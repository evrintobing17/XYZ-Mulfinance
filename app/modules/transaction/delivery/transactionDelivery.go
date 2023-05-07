package delivery

import (
	"github.com/evrintobing17/XYZ-Multifinance/app/helper/jsonresponse"
	"github.com/evrintobing17/XYZ-Multifinance/app/models"
	"github.com/evrintobing17/XYZ-Multifinance/app/modules/transaction"
	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	transactionUC transaction.TransactionUsecase
}

func NewTransactionHTTPHandler(r *gin.Engine, transactionUC transaction.TransactionUsecase) {
	handlers := transactionHandler{
		transactionUC: transactionUC,
	}

	route := r.Group("/v1/transaction")
	{
		route.POST("/", handlers.CreateTransaction)
	}
}

func (handler *transactionHandler) CreateTransaction(c *gin.Context) {

	var transaction models.Transaction
	errBind := c.ShouldBind(&transaction)
	if errBind != nil {
		jsonresponse.BadRequest(c, errBind)
		return
	}

	// Retrieve customer from the database
	err := handler.transactionUC.CreateTransactionWithLock(&transaction)
	if err != nil {
		jsonresponse.BadRequest(c, err)
		return
	}
	jsonresponse.OK(c, "Success")
	return
}
