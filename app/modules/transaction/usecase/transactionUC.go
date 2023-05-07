package usecase

import (
	"errors"
	"fmt"
	"time"

	"github.com/evrintobing17/XYZ-Multifinance/app/models"
	"github.com/evrintobing17/XYZ-Multifinance/app/modules/customer"
	"github.com/evrintobing17/XYZ-Multifinance/app/modules/transaction"
)

type UC struct {
	trxRepo      transaction.TransactionRepo
	customerRepo customer.CustomerRepository
}

func NewTransactionUsecase(trxRepo transaction.TransactionRepo, customerRepo customer.CustomerRepository) transaction.TransactionUsecase {
	return &UC{
		trxRepo:      trxRepo,
		customerRepo: customerRepo,
	}
}

// CreateTransactionWithLock implements transaction.TransactionUsecase
func (uc *UC) CreateTransactionWithLock(transaction *models.Transaction) error {

	customer, err := uc.customerRepo.GetCustomerByID(transaction.CustomerID)
	if err != nil {
		return err
	}

	var totalAmount float64
	for i := range transaction.Limits {
		totalAmount = totalAmount + transaction.Limits[i].Amount
	}

	if totalAmount > customer.Limit {
		fmt.Println()
		return errors.New("Loan exceeding the limit")
	}

	newLimit := customer.Limit - totalAmount

	resultChan := make(chan bool)

	// Execute createTransaction function in a Goroutine
	go func() {
		success, err := uc.trxRepo.CreateTransactionWithLock(transaction)
		if err != nil {
			resultChan <- false // Signal an error
		} else {
			resultChan <- success // Send the transaction status
		}
	}()

	select {
	case success := <-resultChan:
		if !success {
			return errors.New("Failed to create Transaction")
		}
	case <-time.After(15 * time.Second):
		return errors.New("Database operation timed out")

	}

	go uc.customerRepo.UpdateLimit(transaction.CustomerID, newLimit)

	return nil
}
