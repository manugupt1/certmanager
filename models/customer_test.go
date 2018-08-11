package models

import (
	"testing"

	"github.com/gobuffalo/suite"
)

type CustomerSuite struct {
	*suite.Model
}

func Test_Customer(t *testing.T) {
	model := suite.NewModel()
	as := &CustomerSuite{
		Model: model,
	}
	suite.Run(t, as)
}

func (c *CustomerSuite) Test_Validate() {
	expectValidationErrors := []Customer{
		Customer{Name: "M", Email: "manu@example.com", Password: ""},
		Customer{Name: "Manu", Email: "", Password: "Gupta"},
		Customer{Name: "", Email: "manu@example.com", Password: "Gupta"},
		Customer{Name: "M", Email: "manuexample.com", Password: "G"},
	}

	for _, customer := range expectValidationErrors {
		validationErrors, _ := customer.Validate(c.DB)
		if !validationErrors.HasAny() {
			c.Fail("Validation failed while updating, save and create", validationErrors.Error())
		}
	}

	expectToSucceed := []Customer{
		Customer{Name: "M", Email: "manu@example.com", Password: "G"},
	}
	for _, customer := range expectToSucceed {
		validationErrors, _ := customer.Validate(c.DB)

		if validationErrors.HasAny() {
			c.Fail("Validation failed while updating, save and create", validationErrors.Error())
		}
	}
}

func (c *CustomerSuite) Test_Create() {
	expectValidationErrors := []Customer{
		Customer{Name: "M", Email: "manu@example.com", Password: ""},
		Customer{Name: "Manu", Email: "", Password: "Gupta"},
		Customer{Name: "", Email: "manu@example.com", Password: "Gupta"},
		Customer{Name: "M", Email: "manuexample.com", Password: "G"},
	}

	for _, customer := range expectValidationErrors {
		validationErrors, _ := customer.Create(c.DB)
		if !validationErrors.HasAny() {
			c.Fail("Validation failed while updating, save and create", validationErrors.Error())
		}
	}

	expectToSucceed := []*Customer{
		&Customer{Name: "M", Email: "manu@example.com", Password: "G"},
	}

	for _, customer := range expectToSucceed {
		validationErrors, _ := customer.Create(c.DB)

		if validationErrors.HasAny() {
			c.Fail("Validation failed while updating, save and create", validationErrors.Error())
		}
	}
}
