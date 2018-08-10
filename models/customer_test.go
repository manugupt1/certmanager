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
	expectToFail := []Customer{
		Customer{Name: "M", Email: "", Password: ""},
		Customer{Name: "", Email: "E", Password: ""},
		Customer{Name: "", Email: "", Password: "G"},
		Customer{Name: "M", Email: "manuexample.com", Password: "G"},
	}

	for _, customer := range expectToFail {
		validationErrors, _ := customer.Validate(nil)
		if validationErrors.Count() == 0 {
			c.Fail("Validation failed while updating, save and create", validationErrors.Error())
		}
	}

	expectToSucceed := []Customer{
		Customer{Name: "M", Email: "manu@example.com", Password: "G"},
	}
	for _, customer := range expectToSucceed {
		validationErrors, _ := customer.Validate(nil)

		if validationErrors != nil {
			c.Fail("Validation failed while updating, save and create", validationErrors.Error())
		}
	}

}
