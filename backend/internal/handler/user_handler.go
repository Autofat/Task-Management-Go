package handler

import (
	"net/http"
	"strconv"
	"task-management/internal/service"
	"task-management/internal/utils"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) RegisterUser(c *gin.Context){
	var req struct{
		Email string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
		Fullname string `json:"fullname" binding:"required"`
		Role string `json:"role" binding:"omitempty,oneof=ADMIN USER"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	user, err := h.userService.CreateUser(req.Email, req.Password, req.Fullname, req.Role)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Failed to create user", err.Error())
		return
	}

	response := gin.H{
		"id": user.ID,
		"email": user.Email,
		"fullname": user.Fullname,
		"role": user.Role,
	}

	utils.SuccessResponse(c, http.StatusCreated, "User Created Successfully", response)
}

func (h *UserHandler) GetUser(c *gin.Context){
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid user ID", err.Error())
		return
	}
	user, err := h.userService.GetUserByID(uint(id))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "User not found", err.Error())
		return
	}
	
	response := gin.H{
		"id": user.ID,
		"email": user.Email,
		"fullname": user.Fullname,
		"role": user.Role,
	}

	utils.SuccessResponse(c, http.StatusOK, "User Retrieved Successfully", response)
}

func (h *UserHandler) UpdateUser(c *gin.Context){
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid user ID", err.Error())
		return
	}

	var req struct{
		Fullname string `json:"fullname"`
		Role string `json:"role" binding:"omitempty,oneof=ADMIN USER"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	err = h.userService.UpdateUser(uint(id), req.Fullname, req.Role)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Failed to update user", err.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "User Updated Successfully", nil)
}

func (h *UserHandler) DeleteUser(c *gin.Context){
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid user ID", err.Error())
		return
	}
	err = h.userService.DeleteUser(uint(id))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Failed Deteleting User", err.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "User Deleted Successfully", nil)
}