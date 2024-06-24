package middlewares

import (
	"blog/internal/config"
	"blog/internal/constants"
	"blog/internal/utils"
	"errors"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func verifyToken(tokenString string) (bool, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return false, err
	}
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.SecretKey), nil
	})
	if err != nil {
		return false, err
	}
	_, ok := token.Claims.(*jwt.StandardClaims)
	if !ok || !token.Valid {
		return false, errors.New("Token verification failed")
	}
	return true, nil
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader(constants.TOKEN_KEY)
		if token == "" {
			utils.SetCtxResponse(c, nil, http.StatusUnauthorized, "Token must be not empty")
			return
		}

		isValidatedSuccess, err := verifyToken(token)
		if isValidatedSuccess {
			c.Next()
			return
		}
		if err != nil || !isValidatedSuccess {
			utils.SetCtxResponse(c, nil, http.StatusUnauthorized, err.Error())
			return
		}
	}
}
