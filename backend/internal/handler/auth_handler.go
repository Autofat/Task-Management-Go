package handler

import (
	"net/http"
	"task-management/internal/service"
	"task-management/internal/utils"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authSevice *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authSevice: authService,
	}
}
func (h* AuthHandler) Login (c *gin.Context){
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req);  err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	token, userData, err := h.authSevice.Login(req.Email, req.Password)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Login failed", err.Error())
		return
	}
	response := gin.H{
		"token": token,
		"user": userData,
	}
	utils.SuccessResponse(c, http.StatusOK, "Login successful", response)
}

func (h* AuthHandler) Logout (c *gin.Context){
	utils.SuccessResponse(c, http.StatusOK, "Logout successful", nil)
}
