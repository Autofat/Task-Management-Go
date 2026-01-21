package handler

import (
	"net/http"
	"strconv"
	"task-management/internal/service"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	taskService *service.TaskService
}

func NewTaskHandler(taskService *service.TaskService) *TaskHandler {
	return &TaskHandler{
		taskService: taskService,
	}
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Priority	string `json:"priority" binding:"omitempty,oneof=low medium high"`
		ProjectID   uint   `json:"project_id" binding:"required"`
		AssignedID  uint   `json:"assigned_id" binding:"required"`
		DueDate	 	string `json:"due_date"`

	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	task, err := h.taskService.CreateTask(req.Title, req.Description, req.Priority, req.ProjectID, req.AssignedID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Task Created Successfully",
		"data": gin.H{
			"id":          task.ID,
			"title":       task.Title,
			"description": task.Description,
			"project_id":  task.ProjectID,
			"assigned_id": task.AssignedID,
			"status":      task.Status,
			"due_date":    task.DueDate,
		},
	})

}

func (h *TaskHandler) GetTaskByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	projectIDStr := c.Query("project_id")
	projectID, err := strconv.ParseUint(projectIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid project_id",
		})
		return
	}
	task, err := h.taskService.GetTaskByID(uint(id), uint(projectID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"id":          task.ID,
			"title":       task.Title,
			"description": task.Description,
			"project_id":  task.ProjectID,
			"assigned_id": task.AssignedID,
			"status":      task.Status,	
			"due_date":    task.DueDate,
			"priority":    task.Priority,
		},
	})

}

func (h *TaskHandler) GetTasksByProjectID(c *gin.Context) {
	projectIDStr := c.Query("project_id")
	projectID, err := strconv.ParseUint(projectIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid project_id",
		})
		return
	}
	tasks, err := h.taskService.GetTasksByProjectID(uint(projectID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	tasksResponse := []gin.H{}
	for _, task := range tasks {
		tasksResponse = append(tasksResponse, gin.H{
			"id":          task.ID,
			"title":       task.Title,
			"description": task.Description,
			"project_id":  task.ProjectID,
			"assigned_id": task.AssignedID,
			"status":      task.Status,
			"due_date":    task.DueDate,
			"priority":    task.Priority,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": tasksResponse,
	})
}

func (h *TaskHandler) UpdateTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	projectIdStr := c.Query("project_id")
	projectId, err := strconv.ParseUint(projectIdStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid project_id",
		})
		return
	}
	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Status      string `json:"status" binding:"omitempty,oneof=pending in_progress completed"`
		Priority	string `json:"priority" binding:"omitempty,oneof=low medium high"`
		AssignedID  uint   `json:"assigned_id"`
		DueDate	 	string `json:"due_date"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = h.taskService.UpdateTask(
		uint(id), 
		req.Title, 
		req.Description, 
		req.Status, 
		req.DueDate, 
		req.Priority, 
		uint(projectId), 
		req.AssignedID,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Task Updated Successfully",
	})
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	projecrIdStr := c.Query("project_id")
	projectId, err := strconv.ParseUint(projecrIdStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid project_id",
		})
		return
	}
	err = h.taskService.DeleteTask(uint(id), uint(projectId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Task Deleted Successfully",
	})
}