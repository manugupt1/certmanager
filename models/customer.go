package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
)

type Customer struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Name      string    `json:"name" db:"name"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"password" db:"password"`
}

// String is not required by pop and may be deleted
func (c Customer) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Customers is not required by pop and may be deleted
type Customers []Customer

// String is not required by pop and may be deleted
func (c Customers) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Gets a list of all the Customers
func (c *Customers) List(tx *pop.Connection) error {
	return tx.All(c)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// It hashes the password using bcrypt before storing in the database
func (c *Customer) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Name: "Password", Field: c.Password},
		&validators.StringIsPresent{Name: "Name", Field: c.Name},
		&validators.EmailIsPresent{Name: "Email", Field: c.Email, Message: "Email is not in the right format"},
	), nil
}

func (c *Customer) BeforeCreate(tx *pop.Connection) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(c.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.WithStack(err)
	}
	c.Password = string(hash)
	return nil
}

// Create takes a running transaction and tries to add a Customer to the database. If it is not able to create one it will return an appropriate error.
func (c *Customer) Create(tx *pop.Connection) (*validate.Errors, error) {
	return tx.ValidateAndCreate(c)
}
