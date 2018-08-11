package models

import (
	"time"

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
