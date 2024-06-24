package middlewares

import (
	"blog/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func validateParamsError(c *gin.Context, err error) bool {
	validateError, ok := err.(validator.ValidationErrors)
	if ok {
		var errMsg string
		for _, e := range validateError {
			switch e.Tag() {
			case "required":
				errMsg = e.Field() + " is required"
			case "email":
				errMsg = e.Field() + " must be a valid email"
			case "min":
				errMsg = e.Field() + " must be at least " + e.Param() + " characters long"
			}
		}
		utils.SetCtxResponse(c, nil, http.StatusBadRequest, errMsg)
		c.Abort()
	}
	if validateError != nil {
		return true
	}
	return false
}

func Validate(dto interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := c.ShouldBind(dto)
		if isValidatedError := validateParamsError(c, err); isValidatedError {
			return
		}
		if err != nil {
			utils.SetCtxResponse(c, nil, http.StatusBadRequest, err.Error())
			c.Abort()
			return
		}
		utils.SetCtxValidatedData(c, dto)
		c.Next()
	}
}
