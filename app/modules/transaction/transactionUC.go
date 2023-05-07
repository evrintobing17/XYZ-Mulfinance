package transaction

import "github.com/evrintobing17/XYZ-Multifinance/app/models"

type TransactionUsecase interface {
	CreateTransactionWithLock(transaction *models.Transaction) error
}
