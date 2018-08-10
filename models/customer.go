package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"

	"golang.org/x/crypto/bcrypt"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/pkg/errors"
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

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// It hashes the password using bcrypt before storing in the database
func (c *Customer) Validate(tx *pop.Connection) (*validate.Errors, error) {
	validationErrors := validate.Validate(
		&validators.StringIsPresent{Name: "Password", Field: c.Password},
		&validators.StringIsPresent{Name: "Name", Field: c.Name},
		&validators.EmailIsPresent{Name: "Email", Field: c.Email, Message: "Email is not in the right format"},
	)
	if validationErrors.Count() > 0 {
		return validationErrors, nil
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(c.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	c.Password = string(hash)

	return nil, nil
}
