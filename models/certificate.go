package models

import (
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

	err := tx.Where("id::text = ?", id).First(c)
	if err != nil {
		return err
	}
	newCert := c
	newCert.Activated = active
	_, err := tx.ValidateAndUpdate(newCert, "id", "customer_id", "created_at")
	return err
}
