package transaction

import "github.com/evrintobing17/XYZ-Multifinance/app/models"

type TransactionRepo interface {
	CreateTransactionWithLock(transaction *models.Transaction) (bool, error)
}
