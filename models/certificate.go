package models

import (
	"context"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/gobuffalo/uuid"

	"github.com/gobuffalo/pop"
)

const path = "./certificates"

// Certificate is the model that defines certificate and indicates if it is active or not
type Certificate struct {
	ID         int       `json:"id" db:"id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
	Activated  bool      `json:"activated" db:"activated"`
	Customer   Customer  `belongs_to:"customer"`
	CustomerID int       `db:"customer_id"`
}

// Certificates is not required by pop and may be deleted
type Certificates []Certificate

func (c *Certificates) ListCertificate(tx *pop.Connection, customer_id string, active *bool) error {
	var err error
	if active == nil {
		err = tx.Where("customer_id::text = ?", customer_id).All(c)
	} else {
		err = tx.Where("customer_id::text = ? AND activated = ?", customer_id, *active).All(c)
	}

	if err != nil {
		return err
	}
	return nil
}

func (c *Certificate) UpdateStatus(tx *pop.Connection, id int, toActivate bool) error {
	query := `UPDATE certificates SET activated = $1, updated_at = $2 WHERE id=$3`
	_, err := SQL.Exec(query, toActivate, time.Now(), id)
	if err != nil {
		return err
	}
	return nil
}

func (c *Certificate) CreateCertificate(tx *pop.Connection, custID string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start a transaction
	certTx, err := SQL.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// Create a key in the fs
	fsKey, fsBody, fsErr := newCertificate(ctx)

	// If there is a error, call cancel to rollback certTx
	if fsErr != nil {
		cancel()
		return fsErr
	}

	// Add metadata in DB
	dbErr := addCertificateMeta(ctx, custID, fsKey, fsBody)
	if dbErr != nil {
		cancel()
		return dbErr
	}

	commitErr := certTx.Commit()
	// If there is a commit error, rollback everything
	// TODO: Remove dangling certs
	if commitErr != nil {
		return commitErr
	}
	return nil

}

func newCertificate(ctx context.Context) (string, string, error) {
	const cmd = "openssl"
	certName, err := uuid.NewV4()
	if err != nil {
		return "", "", err
	}
	fPath := filepath.Join(path, certName.String())
	opts := []string{"req", "-nodes", "-newkey", "rsa:2048", "-keyout", fPath + ".key", "-out", fPath + ".csr", "-subj", "/C=GB/ST=London/L=London/O=Global Security/OU=IT Department/CN=example.com"}
	cmdObj := exec.Command(cmd, opts...)
	_, err = cmdObj.CombinedOutput()
	if err != nil {
		return "", "", err
	}
	return certName.String() + ".key", certName.String() + ".csr", nil
}

func addCertificateMeta(ctx context.Context, custID string, key, body string) error {
	query := `INSERT INTO certificates (activated, created_at, customer_id, updated_at, key_path, body_path) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := SQL.Exec(query, true, time.Now(), custID, time.Now(), key, body)
	if err != nil {
		return err
	}
	return nil
}
