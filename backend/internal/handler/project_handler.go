package handler

import (
	"net/http"
	"strconv"
	"task-management/internal/service"

	"github.com/gin-gonic/gin"
)

type ProjectHandler struct {
	projectService *service.ProjectService
}

func NewProjectHandler(projectService *service.ProjectService) *ProjectHandler {
	return &ProjectHandler{
		projectService: projectService,
	}
}

func (h *ProjectHandler) CreateProject(c *gin.Context) {
	var req struct {
		Title       string `json:"title"`
		OwnerID     uint   `json:"owner_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	project, err := h.projectService.CreateProject(req.Title, req.OwnerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Project Created Successfully",
		"data": gin.H{
			"id":          project.ID,
			"title":       project.Title,
			"owner_id":    project.OwnerID,
		},
	})

}

func (h *ProjectHandler) GetProjectByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	project, err := h.projectService.GetProjectByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"id":       project.ID,
			"title":    project.Title,
			"owner_id": project.OwnerID,
		},
	})

}

func (h *ProjectHandler) GetProjectsByOwnerID(c *gin.Context) {
	ownerIDStr := c.Query("owner_id")
	if ownerIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "owner_id query parameter is required",
		})
		return
	}
	ownerID, err := strconv.ParseUint(ownerIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	projects, err := h.projectService.GetProjectsByOwnerID(uint(ownerID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
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
	c.JSON(http.StatusOK, gin.H{
		"data": responseProjects,
	})
}

func (h *ProjectHandler) UpdateProject(c *gin.Context) {
	var req struct {
		Title       string `json:"title" `
	}
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = h.projectService.UpdateProject(uint(id), req.Title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Project Updated Successfully",
	})
}

func (h *ProjectHandler) DeleteProject(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = h.projectService.DeleteProject(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Project Deleted Successfully",
	})
}