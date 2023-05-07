package usecase

import (
	"github.com/evrintobing17/XYZ-Multifinance/app/models"
	"github.com/evrintobing17/XYZ-Multifinance/app/modules/customer"
)

type UC struct {
	repo customer.CustomerRepository
}

func NewCustomerUsecase(repo customer.CustomerRepository) customer.CustomerUsecase {
	return &UC{
		repo: repo,
	}
}

// GetCustomerByID implements customer.CustomerUsecase
func (uc *UC) GetCustomerByID(customerID int) (*models.Customer, error) {
	customer, err := uc.repo.GetCustomerByID(customerID)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

// InsertCustomer implements customer.CustomerUsecase
func (uc *UC) InsertCustomer(customer models.Customer) error {
	err := uc.repo.InsertCustomer(customer)
	if err != nil {
		return err
	}
	return nil
}
