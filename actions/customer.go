package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/manugupt1/certmanager/models"
	"github.com/pkg/errors"
)

type CustomerActions struct{}

// Create is the handler that will create a new user if the user does not exist and validations pass
func (cr CustomerActions) Create(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	if tx == nil {
		err := "Database connection lost"
		return c.Render(500, r.JSON(&err))
	}

	customer := &models.Customer{}

	if err := c.Bind(customer); err != nil {
		return errors.WithStack(err)
	}

	validationErrors, err := customer.Create(tx)
	if err != nil {
		return err
	}

	if validationErrors.HasAny() {
		return c.Render(http.StatusBadRequest, r.JSON(validationErrors.Errors))
	}

	return c.Render(http.StatusOK, nil)
}

// List is the handler that will list all the customers stored in the db
func (cr CustomerActions) List(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	if tx == nil {
		err := "Database connection lost"
		return c.Render(500, r.JSON(&err))
	}
	customers := &models.Customers{}
	customers.List(tx)
	return c.Render(http.StatusOK, r.JSON(customers))
}

// Delete is the handler which will delete a customer from the model
func (cr CustomerActions) Delete(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	if tx == nil {
		err := "Database connection lost"
		return c.Render(500, r.JSON(&err))
	}

	customer := &models.Customer{}
	if err := c.Bind(customer); err != nil {
		return errors.WithStack(err)
	}

	err := customer.Delete(tx)
	if err != nil {
		errMsg := "No records found"
		return c.Render(http.StatusBadRequest, r.JSON(&errMsg))
	}
	return c.Render(http.StatusOK, nil)
}
