package repository

import (
	"database/sql"

	"github.com/evrintobing17/XYZ-Multifinance/app/models"
	"github.com/evrintobing17/XYZ-Multifinance/app/modules/customer"
)

type repository struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) customer.CustomerRepository {

	return &repository{
		db: db,
	}

}

// GetCustomerByID implements customer.CustomerRepository
func (r *repository) GetCustomerByID(customerID int) (*models.Customer, error) {
	query := "SELECT id, nik, full_name, legal_name, birthplace, birthdate, salary, ktp_photo, selfie_photo, limits FROM customers WHERE id = ?"
	row := r.db.QueryRow(query, customerID)

	// Initialize a Customer struct to hold the retrieved data
	customer := &models.Customer{}
	err := row.Scan(&customer.ID, &customer.NIK, &customer.FullName, &customer.LegalName, &customer.Birthplace, &customer.Birthdate, &customer.Salary, &customer.KTPPhoto, &customer.SelfiePhoto, &customer.Limit)
	if err != nil {
		return nil, err
	}

	// Retrieve limits for the customer

	return customer, nil
}

// GetLimitByCustomerID implements customer.CustomerRepository
func (r *repository) GetLimitByCustomerID(customerID int) ([]models.Limit, error) {
	query := "SELECT tenor, amount FROM limits WHERE customer_id = ?"
	rows, err := r.db.Query(query, customerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Initialize a slice to hold the retrieved limits
	limits := []models.Limit{}

	// Iterate over the rows and populate the limits slice
	for rows.Next() {
		var limit models.Limit
		if err := rows.Scan(&limit.Tenor, &limit.Amount); err != nil {
			return nil, err
		}
		limits = append(limits, limit)
	}

	return limits, nil
}

// InsertCustomer implements customer.CustomerRepository
func (r *repository) InsertCustomer(customer models.Customer) error {
	query := "INSERT INTO customers (nik, full_name, legal_name, birthplace, birthdate, salary, ktp_photo, selfie_photo, limits) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, err := r.db.Exec(query, customer.NIK, customer.FullName, customer.LegalName, customer.Birthplace, customer.Birthdate, customer.Salary, customer.KTPPhoto, customer.SelfiePhoto, customer.Limit)
	if err != nil {
		return err
	}

	return nil
}
func (r *repository) UpdateLimit(id int, limit float64) error {
	query := "UPDATE customers SET limits = ? WHERE id = ?"
	_, err := r.db.Exec(query, limit, id)
	if err != nil {
		return err
	}

	return nil
}
