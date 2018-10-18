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

func toByte(str string) []byte {
	return []byte(str)
}
func (c *CustomerSuite) Test_Validate() {
	expectValidationErrors := []Customer{
		Customer{Name: "M", Email: "manu@example.com", Password: toByte("")},
		Customer{Name: "Manu", Email: "", Password: toByte("Gupta")},
		Customer{Name: "", Email: "manu@example.com", Password: toByte("Gupta")},
		Customer{Name: "M", Email: "manuexample.com", Password: toByte("G")},
	}

	for _, customer := range expectValidationErrors {
		validationErrors, _ := customer.Validate(c.DB)
		if !validationErrors.HasAny() {
			c.Fail("Validation failed while updating, save and create", validationErrors.Error())
		}
	}

	expectToSucceed := []Customer{
		Customer{Name: "M", Email: "manu@example.com", Password: toByte("G")},
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
		Customer{Name: "M", Email: "manu@example.com", Password: toByte("")},
		Customer{Name: "Manu", Email: "", Password: toByte("Gupta")},
		Customer{Name: "", Email: "manu@example.com", Password: toByte("Gupta")},
		Customer{Name: "M", Email: "manuexample.com", Password: toByte("G")},
	}

	for _, customer := range expectValidationErrors {
		validationErrors, _ := customer.Create(c.DB)
		if !validationErrors.HasAny() {
			c.Fail("Validation failed while updating, save and create", validationErrors.Error())
		}
	}

	expectToSucceed := []*Customer{
		&Customer{Name: "M", Email: "manu@example.com", Password: toByte("G")},
	}

	for _, customer := range expectToSucceed {
		validationErrors, _ := customer.Create(c.DB)

		if validationErrors.HasAny() {
			c.Fail("Validation failed while updating, save and create", validationErrors.Error())
		}
	}

	duplicateEmail := &Customer{Name: "m", Email: "manu@example.com", Password: toByte("G")}
	vErr, _ := duplicateEmail.Create(c.DB)
	if !vErr.HasAny() {
		c.Fail("Expecting to fail on creating a customer with duplicate email")
	}
}
