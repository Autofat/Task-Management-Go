package handler

import (
	"net/http"
	"strconv"
	"task-management/internal/service"
	"task-management/internal/utils"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	taskService *service.TaskService
	projectMemberService *service.ProjectMemberService
}

func NewTaskHandler(taskService *service.TaskService, projectMemberService *service.ProjectMemberService) *TaskHandler {
	return &TaskHandler{
		taskService: taskService,
		projectMemberService: projectMemberService,
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

	userId := c.GetUint("user_id")

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	isMember, err := h.projectMemberService.IsMember(req.ProjectID, userId)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to verify project membership", err.Error())
		return
	}
	if !isMember {
		utils.ErrorResponse(c, http.StatusForbidden, "You are not a member of this project", nil)
		return
	}

	task, err := h.taskService.CreateTask(req.Title, req.Description, req.Priority, req.ProjectID, req.AssignedID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Failed to create task", err.Error())
		return
	}

	response := gin.H{
			"id":          task.ID,
			"title":       task.Title,
			"description": task.Description,
			"project_id":  task.ProjectID,
			"assigned_id": task.AssignedID,
			"status":      task.Status,	
			"due_date":    task.DueDate,
			"priority":    task.Priority,
	}

	utils.SuccessResponse(c, http.StatusCreated, "Task Created Successfully", response)
}

func (h *TaskHandler) GetTaskByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "invalid task ID", err.Error())
		return
	}
	
	isMember, err := h.projectMemberService.IsMember(uint(id), c.GetUint("user_id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to verify project membership", err.Error())
		return
	}
	if !isMember {
		utils.ErrorResponse(c, http.StatusForbidden, "Access Denied", nil)
		return
	}

	projectIDStr := c.Query("project_id")
	projectID, err := strconv.ParseUint(projectIDStr, 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "invalid project ID", err.Error())
		return
	}
	task, err := h.taskService.GetTaskByID(uint(id), uint(projectID))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Task not found", err.Error())
		return
	}

	response := gin.H{
			"id":          task.ID,
			"title":       task.Title,
			"description": task.Description,
			"project_id":  task.ProjectID,
			"assigned_id": task.AssignedID,
			"status":      task.Status,	
			"due_date":    task.DueDate,
			"priority":    task.Priority,
		}
	utils.SuccessResponse(c, http.StatusOK, "Task Retrieved Successfully", response)
}

func (h *TaskHandler) GetTasksByProjectID(c *gin.Context) {
	projectIDStr := c.Query("project_id")
	projectID, err := strconv.ParseUint(projectIDStr, 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid project ID", err.Error())
		return
	}
	tasks, err := h.taskService.GetTasksByProjectID(uint(projectID))
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve tasks", err.Error())
		return
	}

	isMember, err := h.projectMemberService.IsMember(uint(projectID), c.GetUint("user_id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to verify project membership", err.Error())
		return
	}
	if !isMember {
		utils.ErrorResponse(c, http.StatusForbidden, "Access Denied", nil)
		return
	}

	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")
	sort := c.DefaultQuery("sort", "created_at")
	order := c.DefaultQuery("order", "asc")
	status := c.Query("status")
	priority := c.Query("priority")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid page parameter", nil)
		return
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid limit parameter", nil)
		return
	}

	validStatuses := map[string]bool{"pending": true, "in_progress": true, "completed": true}
	if status != "" && !validStatuses[status] {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid status parameter", nil)
		return
	}
	validPriorities := map[string]bool{"low": true, "medium": true, "high": true}
	if priority != "" && !validPriorities[priority] {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid priority parameter", nil)
		return
	}

	tasks, total, err := h.taskService.GetTasksByProjectIDWithFilters(uint(projectID), status, priority, page, limit, sort, order )
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve tasks with filters", err.Error())
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

	response := gin.H{
		"data": 	 tasksResponse,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
			"pages": (total + int64(limit) -1) / int64(limit),
		},
	}

	utils.SuccessResponse(c, http.StatusOK, "Tasks Retrieved Successfully", response)
}



func (h *TaskHandler) UpdateTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid Task ID", err.Error())
		return
	}

	projectIdStr := c.Query("project_id")
	projectId, err := strconv.ParseUint(projectIdStr, 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid project ID", err.Error())
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
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	isMember, err := h.projectMemberService.IsMember(uint(projectId), c.GetUint("user_id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to verify project membership", err.Error())
		return
	}
	if !isMember {
		utils.ErrorResponse(c, http.StatusForbidden, "Access Denied", nil)
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
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update task", err.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Task Updated Successfully", nil)
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid Task ID", err.Error())
		return
	}

	projectIdStr := c.Query("project_id")
	projectId, err := strconv.ParseUint(projectIdStr, 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid Project ID", err.Error())
		return
	}
	
	isMember, err := h.projectMemberService.IsMember(uint(projectId), c.GetUint("user_id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to verify project membership", err.Error())
		return
	}
	if !isMember {
		utils.ErrorResponse(c, http.StatusForbidden, "Access Denied", nil)
		return
	}
	
	err = h.taskService.DeleteTask(uint(id), uint(projectId))
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError,"Failed Deleting Task", err.Error())
		return
	}
	
	utils.SuccessResponse(c, http.StatusOK, "Task Deleted Successfully", nil)
}