package models

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
)

// Certificate is the model that defines certificate and indicates if it is active or not
type Certificate struct {
	ID         uuid.UUID `json:"id" db:"id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
	Activated  bool      `json:"activated" db:"activated"`
	Customer   Customer  `belongs_to:"customer"`
	CustomerID uuid.UUID `db:"customer_id"`
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

func (c *Certificate) UpdateStatus(tx *pop.Connection, id string, active bool) error {

	err := tx.Find(c, "068a27c2-1851-4367-aa51-813e7644ec23")
	if err != nil {
		return err
	}

	c.Activated = active
	v, err := tx.ValidateAndUpdate(c, "id", "customer_id", "created_at")
	if v.HasAny() {
		fmt.Println(v.Errors)
	}
	return err
}

func (c *Certificate) Create(tx *pop.Connection, cust_id string) error {
	// customer := &Customer{}
	// err := customer.Find(tx, cust_id)
	// if err != nil {
	// 	return err
	// }

	certID, err := uuid.NewV4()
	if err != nil {
		return err
	}
	// c.ID = certID
	// c.CustomerID = customer.ID
	// c.Activated = true
	// v, err := tx.ValidateAndCreate(c)

	tx.RawQuery("INSERT INTO certificates (id, customer_id, created_at, updated_at, activated) VALUES(?, ?, ?, ?, ?)", certID, "068a27c2-1851-4367-aa51-813e7644ecb7", time.Now(), time.Now(), false).All(&c)
	// fmt.Println(v.Erroxrs, err)
	// newCertificate(c.ID.String())
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
