package actions

import (
	"fmt"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/manugupt1/certmanager/models"
	"github.com/pkg/errors"
)

// import (
// 	"github.com/gobuffalo/buffalo"
// 	"github.com/manugupt1/certmanager/models"
// )

type CustomerActions struct{}

func (cr CustomerActions) Create(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	customer := &models.Customer{}

	if err := c.Bind(customer); err != nil {
		return errors.WithStack(err)
	}

	validationErrors, err := customer.Create(tx)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	if validationErrors.HasAny() {
		return c.Render(422, r.JSON(validationErrors.Errors))
	}

	return c.Render(200, nil)
}

func (cr CustomerActions) List(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	customers := &models.Customers{}
	customers.List(tx)
	return c.Render(200, r.JSON(customers))
}
