package actions

import (
	"github.com/manugupt1/certmanager/models"
)

func (as *ActionSuite) Test_CreateCustomer() {
	expectValidationErrors := []models.Customer{
		models.Customer{Name: "M", Email: "manu@example.com", Password: ""},
		models.Customer{Name: "Manu", Email: "", Password: "Gupta"},
		models.Customer{Name: "", Email: "manu@example.com", Password: "Gupta"},
		models.Customer{Name: "M", Email: "manuexample.com", Password: "G"},
	}

	for _, customer := range expectValidationErrors {
		res := as.JSON("/customer").Post(customer)
		as.Equal(422, res.Code)
	}

	expectSuccess := []models.Customer{
		models.Customer{Name: "M", Email: "manu3@example.com", Password: "password"},
		models.Customer{Name: "M", Email: "manu4@example.com", Password: "password"},
	}

	for _, customer := range expectSuccess {
		res := as.JSON("/customer").Post(customer)
		as.Equal(200, res.Code, res.Body.String())
		as.Empty(res.Body)
	}

	duplicateEmail := []models.Customer{
		models.Customer{Name: "M", Email: "manu4@example.com", Password: "password"},
	}
	res := as.JSON("/customer").Post(duplicateEmail[0])
	as.Equal(422, res.Code, res.Body.String())
}

//Test_CustomerList tests getting a list of customers
func (as *ActionSuite) Test_ListCustomer() {
	res := as.JSON("/customer").Get()
	as.Equal(200, res.Code)
}
