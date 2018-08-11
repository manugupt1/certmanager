package actions

import "github.com/manugupt1/certmanager/models"

func (as *ActionSuite) Test_CustomerCreateHandler() {
	expectValidationErrors := []models.Customer{
		models.Customer{Name: "M", Email: "manu@example.com", Password: ""},
		models.Customer{Name: "Manu", Email: "", Password: "Gupta"},
		models.Customer{Name: "", Email: "manu@example.com", Password: "Gupta"},
		models.Customer{Name: "M", Email: "manuexample.com", Password: "G"},
	}

	for _, customer := range expectValidationErrors {
		res := as.JSON("/customer/create").Post(customer)
		as.Equal(422, res.Code)
	}

	expectSuccess := []models.Customer{
		models.Customer{Name: "M", Email: "manu3@example.com", Password: "password"},
		models.Customer{Name: "M", Email: "manu4@example.com", Password: "password"},
	}

	for _, customer := range expectSuccess {
		res := as.JSON("/customer/create").Post(customer)
		as.Equal(200, res.Code, res.Body.String())
		as.Empty(res.Body)
	}

	// TODO: Duplicate email fails but the response should be 422
	// With an appropriate error messages
	duplicateEmail := []models.Customer{
		models.Customer{Name: "M", Email: "manu4@example.com", Password: "password"},
	}
	res := as.JSON("/customer/create").Post(duplicateEmail[0])
	as.Equal(422, res.Code, res.Body.String())
}

func (as *ActionSuite) Test_CustomerList() {
	res := as.JSON("/customer/list").Get()
	as.Equal(200, res.Code)
}
