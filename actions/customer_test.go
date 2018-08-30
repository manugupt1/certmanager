package actions

import (
	"net/http"

	"github.com/manugupt1/certmanager/models"
)

func toByte(str string) []byte {
	return []byte(str)
}

func (as *ActionSuite) Test_CreateCustomer() {
	expectValidationErrors := []models.Customer{
		models.Customer{Name: "M", Email: "manu@example.com", Password: toByte("")},
		models.Customer{Name: "Manu", Email: "", Password: toByte("Gupta")},
		models.Customer{Name: "", Email: "manu@example.com", Password: toByte("Gupta")},
		models.Customer{Name: "M", Email: "manuexample.com", Password: toByte("G")},
	}

	for _, customer := range expectValidationErrors {
		res := as.JSON("/customer").Post(customer)
		as.Equal(http.StatusBadRequest, res.Code)
	}

	expectSuccess := []models.Customer{
		models.Customer{Name: "M", Email: "manu3@example.com", Password: toByte("password")},
		models.Customer{Name: "M", Email: "manu4@example.com", Password: toByte("password")},
	}

	for _, customer := range expectSuccess {
		res := as.JSON("/customer").Post(customer)
		as.Equal(http.StatusOK, res.Code, res.Body.String())
		as.Empty(res.Body)
	}

	duplicateEmail := []models.Customer{
		models.Customer{Name: "M", Email: "manu4@example.com", Password: toByte("password")},
	}
	res := as.JSON("/customer").Post(duplicateEmail[0])
	as.Equal(http.StatusBadRequest, res.Code, res.Body.String())
}

//Test_CustomerList tests getting a list of customers
func (as *ActionSuite) Test_ListCustomer() {
	res := as.JSON("/customer").Get()
	as.Equal(http.StatusOK, res.Code)
}
