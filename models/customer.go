package models

import (
	"time"

	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
)

// Customer model is the reflection of table column 'customers' in the database
// It implements queries that return single result or when results are to be created / updated / deleted.
type Customer struct {
	ID           uuid.UUID    `json:"id" db:"id"`
	CreatedAt    time.Time    `db:"created_at"`
	UpdatedAt    time.Time    `db:"updated_at"`
	Name         string       `json:"name,omitempty" db:"name"`
	Email        string       `json:"email,omitempty" db:"email"`
	Password     string       `json:"password,omitempty" db:"password"`
	Certificates Certificates `json:"certificates,omitempty" has_many:"certificates" order_by:"created_at desc"`
}

// Customers implements queries that can return more than 1 result from the model
type Customers []Customer

// EmailNotTaken is a custom validator to make sure that a user is not trying to insert the same email twice
type EmailNotTaken struct {
	Name  string
	Field string
	tx    *pop.Connection
}

// IsValid actually tries to check if the email is unique or not
func (v *EmailNotTaken) IsValid(errors *validate.Errors) {
	query := v.tx.Where("email = ?", v.Field)
	queryEmail := &Customer{}
	err := query.First(queryEmail)
	if err == nil {
		errors.Add(validators.GenerateKey(v.Name), "A customer with the email already exists")
	}
}

// List returns all the customers
func (c *Customers) List(tx *pop.Connection) error {
	return tx.Select("id", "name", "email").All(c)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// It hashes the password using bcrypt before storing in the database
func (c *Customer) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Name: "Password", Field: c.Password},
		&validators.StringIsPresent{Name: "Name", Field: c.Name},
		&validators.EmailIsPresent{Name: "Email", Field: c.Email, Message: "Email is not in the right format"},
		&EmailNotTaken{Name: "Email", Field: c.Email, tx: tx},
	), nil
}

func (c *Customer) BeforeCreate(tx *pop.Connection) error {
	id, err := uuid.NewV4()
	if err != nil {
		return errors.WithStack(err)
	}
	c.ID = id
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

func (c *Customer) Delete(tx *pop.Connection) error {
	err := tx.Where("email = (?)", c.Email).First(c)
	if err != nil {
		return err
	}
	tx.Destroy(c)
	return nil
}

func (c *Customer) ListCertificate(tx *pop.Connection) error {
	err := tx.Load(c, "Certificates")
	// err := tx.RawQuery("select * from customers inner join certificates ON customers.id = certificates.customer_id").All(c)
	if err != nil {
		return err
	}
	return nil
}
