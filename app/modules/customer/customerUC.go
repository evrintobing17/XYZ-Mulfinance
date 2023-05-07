package customer

import "github.com/evrintobing17/XYZ-Multifinance/app/models"

type CustomerUsecase interface {
	GetCustomerByID(customerID int) (*models.Customer, error)
	InsertCustomer(customer models.Customer) error
}
