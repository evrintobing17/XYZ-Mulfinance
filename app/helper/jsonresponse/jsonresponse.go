package jsonresponse

import (
	"net/http"

	"github.com/evrintobing17/XYZ-Multifinance/app/models"
	"github.com/gin-gonic/gin"
)

func OK(c *gin.Context, payload interface{}) {
	c.JSON(http.StatusOK, models.Response{
		Message: payload,
	})
}

func BadRequest(c *gin.Context, payload interface{}) {
	c.JSON(http.StatusBadRequest, models.Response{
		Message: payload,
	})
}
