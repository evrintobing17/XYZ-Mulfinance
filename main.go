package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	customerdelivery "github.com/evrintobing17/XYZ-Multifinance/app/modules/customer/delivery"
	customerrepository "github.com/evrintobing17/XYZ-Multifinance/app/modules/customer/repository"
	customerusecase "github.com/evrintobing17/XYZ-Multifinance/app/modules/customer/usecase"
	"github.com/joho/godotenv"

	transactiondelivery "github.com/evrintobing17/XYZ-Multifinance/app/modules/transaction/delivery"
	transactionrepository "github.com/evrintobing17/XYZ-Multifinance/app/modules/transaction/repository"
	transactionusecase "github.com/evrintobing17/XYZ-Multifinance/app/modules/transaction/usecase"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

type Config struct {
	Port string
	DB   *sql.DB
}

func main() {
	// Initialize database connection
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error opening env, %v", err)
	} else {
		fmt.Println(".env file loaded")
	}
	cfg := &mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               "",
		Addr:                 os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
		DBName:               os.Getenv("DB_NAME"),
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// Initialize configuration
	config := &Config{
		Port: os.Getenv("PORT"),
		DB:   db,
	}

	// Initialize router
	r := gin.New()

	// Initialize routes

	//repository
	customerRepo := customerrepository.NewCustomerRepository(db)
	trxRepo := transactionrepository.NewTransactionRepository(db)

	//Usecase
	customerUC := customerusecase.NewCustomerUsecase(customerRepo)
	transactionUC := transactionusecase.NewTransactionUsecase(trxRepo, customerRepo)

	//delivery
	customerdelivery.NewCustomerHTTPHandler(r, customerUC)
	transactiondelivery.NewTransactionHTTPHandler(r, transactionUC)

	// Start server
	log.Printf("Starting server on port %v", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, r))
}
