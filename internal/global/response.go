package global

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseOk(c *gin.Context, data interface{}, message string) {
	c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: message,
		Data:    data,
	})
}

func ResponseError(c *gin.Context, code int, err error) {
	c.JSON(http.StatusOK, Response{
		Code:    http.StatusUnauthorized,
		Message: err.Error(),
		Data:    nil,
	})
}
