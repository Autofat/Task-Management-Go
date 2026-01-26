package handler

import (
	"net/http"
	"strconv"
	"task-management/internal/service"
	"task-management/internal/utils"

	"github.com/gin-gonic/gin"
)

type ProjectHandler struct {
	projectService *service.ProjectService
	projectMemberService *service.ProjectMemberService
}

func NewProjectHandler(projectService *service.ProjectService, projectMemberService *service.ProjectMemberService) *ProjectHandler {
	return &ProjectHandler{
		projectService: projectService,
		projectMemberService: projectMemberService,
	}
}

func (h *ProjectHandler) CreateProject(c *gin.Context) {
	var req struct {
		Title       string `json:"title"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	ownerID := c.GetUint("user_id")

	project, err := h.projectService.CreateProject(req.Title, ownerID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create project", err.Error())
		return
	}
	response := gin.H{
		"id":          project.ID,
		"title":       project.Title,
		"owner_id":    project.OwnerID,
	}
	utils.SuccessResponse(c, http.StatusCreated, "Project Created Successfully", response)
}

func (h *ProjectHandler) GetProjectByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid project ID", err.Error())
		return
	}

	project, err := h.projectService.GetProjectByID(uint(id))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Project not found", err.Error())
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
	
	response := gin.H{
		"id":       project.ID,
		"title":    project.Title,
		"owner_id": project.OwnerID,
	}

	utils.SuccessResponse(c, http.StatusOK, "Project Retrieved Successfully", response)

}

func (h *ProjectHandler) GetProjectsByOwnerID(c *gin.Context) {

	ownerID := c.GetUint("user_id")
	projects, err := h.projectService.GetProjectsByOwnerID(ownerID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve projects", err.Error())
		return
	}

	var responseProjects []gin.H
	for _, project := range projects {
		responseProjects = append(responseProjects, gin.H{
			"id":       project.ID,
			"title":    project.Title,
			"owner_id": project.OwnerID,
			"owner": gin.H{
				"id":       project.Owner.ID,
				"email":    project.Owner.Email,
				"fullname": project.Owner.Fullname,
			},
		})
	}
	utils.SuccessResponse(c, http.StatusOK, "Projects Retrieved Successfully", responseProjects)
}

func (h *ProjectHandler) UpdateProject(c *gin.Context) {
	var req struct {
		Title       string `json:"title"`
	}
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid project ID", err.Error())
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	project, err := h.projectService.GetProjectByID(uint(id))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Project not found", err.Error())
		return
	}

	userID := c.GetUint("user_id")
	if project.OwnerID != userID {
		utils.ErrorResponse(c, http.StatusForbidden, "Access Denied", nil)
		return
	}

	err = h.projectService.UpdateProject(uint(id), req.Title)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update project", err.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Project Updated Successfully", nil)
}

func (h *ProjectHandler) DeleteProject(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid project ID", err.Error())
		return
	}
	
	project, err := h.projectService.GetProjectByID(uint(id))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Project not found", err.Error())
		return
	}
	
	userID := c.GetUint("user_id")
	if project.OwnerID != userID {
		utils.ErrorResponse(c, http.StatusForbidden, "Access Denied", nil)
		return
	}

	err = h.projectService.DeleteProject(uint(id))
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete project", err.Error())
		return
	}
	
	utils.SuccessResponse(c, http.StatusOK, "Project Deleted Successfully", nil)
}