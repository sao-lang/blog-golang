package middlewares

import (
	"net/http"

	"blog/internal/utils"

	"github.com/gin-gonic/gin"
)

type response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Response() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		code, exists := utils.GetCtxResponseStatusCode(c)
		if !exists {
			c.JSON(code, response{
				Code:    http.StatusInternalServerError,
				Message: "internal server error",
				Data:    nil,
			})
			return
		}

		message, exists := utils.GetCtxResponseMessage(c)
		if !exists {
			c.JSON(code, response{
				Code:    http.StatusInternalServerError,
				Message: "internal server error",
				Data:    nil,
			})
			return
		}
		data, exists := utils.GetCtxResponseData(c)
		if !exists {
			data = nil
		}
		c.JSON(code, response{
			Code:    code,
			Message: message,
			Data:    data,
		})
	}
}
