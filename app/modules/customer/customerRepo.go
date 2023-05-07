package customer

import (
	"github.com/evrintobing17/XYZ-Multifinance/app/models"
)

type CustomerRepository interface {
	GetCustomerByID(customerID int) (*models.Customer, error)
	GetLimitByCustomerID(customerID int) ([]models.Limit, error)
	InsertCustomer(customer models.Customer) error
	UpdateLimit(id int, limit float64) error
}
