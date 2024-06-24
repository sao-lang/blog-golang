package handlers

import (
	"blog/internal/dto"
	"blog/internal/services"
	"blog/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *services.UserService
}

func NewAuthHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) Register(c *gin.Context) {
	userParams, _ := utils.GetCtxValidatedData(c)
	if user, err := h.userService.Register(userParams.(*dto.CreateUserDTO)); err != nil {
		utils.SetCtxResponse(c, user, http.StatusFound, err.Error())
		return
	}
	utils.SetCtxResponse(c, nil, http.StatusCreated, "Registered successfully!")
}

func (h *UserHandler) Login(c *gin.Context) {

	userParams, _ := utils.GetCtxValidatedData(c)
	_, token, err := h.userService.Authenticate(userParams.(*dto.CreateUserDTO).UserName, userParams.(*dto.CreateUserDTO).Password)

	if err != nil {
		utils.SetCtxResponse(c, nil, http.StatusUnauthorized, err.Error())
		return
	}
	utils.SetCtxResponse(c, gin.H{"token": token}, http.StatusOK, "Login successfully!")
}
