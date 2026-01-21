package handler

import (
	"net/http"
	"strconv"
	"task-management/internal/service"

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
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := h.userService.CreateUser(req.Email, req.Password, req.Fullname, req.Role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User Created Successfully",
		"data" : gin.H{
			"id": user.ID,
			"email": user.Email,
			"fullname": user.Fullname,
			"role": user.Role,
		},
	})
}

func (h *UserHandler) GetUser(c *gin.Context){
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user, err := h.userService.GetUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User Retrieved Successfully",
		"data": gin.H{
			"id": user.ID,
			"email": user.Email,
			"fullname": user.Fullname,
			"role": user.Role,
		},
	})
}

func (h *UserHandler) UpdateUser(c *gin.Context){
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var req struct{
		Fullname string `json:"fullname"`
		Role string `json:"role" binding:"omitempty,oneof=ADMIN USER"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = h.userService.UpdateUser(uint(id), req.Fullname, req.Role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User Updated Successfully",
	})
}

func (h *UserHandler) DeleteUser(c *gin.Context){
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = h.userService.DeleteUser(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User Deleted Successfully",
	})
}