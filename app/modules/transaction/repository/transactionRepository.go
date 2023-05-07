package repository

import (
	"database/sql"

	"github.com/evrintobing17/XYZ-Multifinance/app/models"
	"github.com/evrintobing17/XYZ-Multifinance/app/modules/transaction"
)

type repository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) transaction.TransactionRepo {

	return &repository{
		db: db,
	}

}

func (r *repository) CreateTransactionWithLock(transaction *models.Transaction) (bool, error) {
	// Begin a transaction
	tx, err := r.db.Begin()
	if err != nil {
		return false, err
	}
	defer tx.Rollback()

	// Acquire a lock on the customer record
	lockStmt := "SELECT * FROM customers WHERE id = ? FOR UPDATE"
	_, err = tx.Exec(lockStmt, transaction.CustomerID)
	if err != nil {
		return false, err
	}

	insertStmt := "INSERT INTO transactions (contract_number, otr, admin_fee, installments, interest, asset_name, customer_id) VALUES (?, ?, ?, ?, ?, ?, ?)"
	_, err = tx.Exec(insertStmt, transaction.ContractNo, transaction.OTR, transaction.AdminFee, transaction.Installment, transaction.Interest, transaction.AssetName, transaction.CustomerID)
	if err != nil {
		return false, err
	}

	for _, limit := range transaction.Limits {
		query := "INSERT INTO limits (customer_id, tenor, amount) VALUES (?, ?, ?)"
		_, err := tx.Exec(query, transaction.CustomerID, limit.Tenor, limit.Amount)
		if err != nil {
			return false, err
		}
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return false, err
	}

	return true, nil
}
