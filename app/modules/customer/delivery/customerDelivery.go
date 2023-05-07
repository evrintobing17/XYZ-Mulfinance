package delivery

import (
	"strconv"

	"github.com/evrintobing17/XYZ-Multifinance/app/helper/jsonresponse"
	"github.com/evrintobing17/XYZ-Multifinance/app/models"
	"github.com/evrintobing17/XYZ-Multifinance/app/modules/customer"
	"github.com/gin-gonic/gin"
)

type customerHandler struct {
	customerUC customer.CustomerUsecase
}

func NewCustomerHTTPHandler(r *gin.Engine, customerUC customer.CustomerUsecase) {
	handlers := customerHandler{
		customerUC: customerUC,
	}

	route := r.Group("/v1/customer")
	{
		route.GET("/:customerID", handlers.GetCustomer)
		route.POST("/", handlers.CreateCustomer)
	}
}

func (handler *customerHandler) GetCustomer(c *gin.Context) {

	customerID, err := strconv.Atoi(c.Param("customerID"))
	if err != nil {
		jsonresponse.BadRequest(c, err)
		return
	}

	// Retrieve customer from the database
	customer, err := handler.customerUC.GetCustomerByID(customerID)
	if err != nil {
		jsonresponse.BadRequest(c, err)
		return
	}
	jsonresponse.OK(c, customer)
	return
}

func (handler *customerHandler) CreateCustomer(c *gin.Context) {

	var customer models.Customer
	errBind := c.ShouldBind(&customer)
	if errBind != nil {
		jsonresponse.BadRequest(c, errBind)
		return
	}

	// Retrieve customer from the database
	err := handler.customerUC.InsertCustomer(customer)
	if err != nil {
		jsonresponse.BadRequest(c, err)
		return
	}
	jsonresponse.OK(c, "Success")
	return
}
