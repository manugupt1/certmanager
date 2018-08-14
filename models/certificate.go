package models

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/gobuffalo/pop"
)

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
	r, err := SQL.Exec(query, toActivate, time.Now(), id)
	fmt.Println("here", r, err)
	if err != nil {
		return err
	}
	return nil
}

func (c *Certificate) CreateCertificate(tx *pop.Connection, cust_id string) error {
	query := `INSERT INTO certificates (activated, created_at, customer_id, updated_at) VALUES ($1, $2, $3, $4)`
	_, err := SQL.Exec(query, true, time.Now(), 1, time.Now())
	if err != nil {
		return err
	}
	return nil
}

func newCertificate(id string) error {
	const path = "./certificates"
	const cmd = "openssl"
	id = filepath.Join(path, id)
	opts := []string{"req", "-nodes", "-newkey", "rsa:2048", "-keyout", id + ".key", "-out", id + ".csr", "-subj", "/C=GB/ST=London/L=London/O=Global Security/OU=IT Department/CN=example.com"}
	cmdObj := exec.Command(cmd, opts...)
	_, err := cmdObj.CombinedOutput()
	if err != nil {
		return err
	}
	return nil
}
