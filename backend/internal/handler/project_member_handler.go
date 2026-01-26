package handler

import (
	"net/http"
	"strconv"
	"task-management/internal/service"
	"task-management/internal/utils"

	"github.com/gin-gonic/gin"
)

type ProjectMemberHandler struct {
	projectMemberService *service.ProjectMemberService
}

func NewProjectMemberHandler(projectMemberService *service.ProjectMemberService) *ProjectMemberHandler {
	return &ProjectMemberHandler{
		projectMemberService: projectMemberService,
	}
}


func (h *ProjectMemberHandler) InviteMember(c *gin.Context) {
	
	// HARD CODE FOR JWT
	inviterID := c.GetUint("user_id")
	
	projectIDstr := c.Param("id")
	projectID, err := strconv.ParseUint(projectIDstr, 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid project ID", err.Error())
		return
	}

	var req struct {
		UserID uint `json:"user_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	err = h.projectMemberService.InviteMember(uint(projectID), req.UserID, inviterID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusForbidden, "Failed to invite member", err.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusCreated, "Member invited successfully", nil)

}

func (h *ProjectMemberHandler) GetProjectMembers(c *gin.Context) {
	projectIDstr := c.Param("id")
	projectID, err := strconv.ParseUint(projectIDstr, 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid project ID", err.Error())
		return
	}

	members, err := h.projectMemberService.GetProjectMembers(uint(projectID))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Failed to get project members", err.Error())
		return
	}

	var memberResponses []gin.H
	for _, member := range members {
		memberResponses = append(memberResponses, gin.H{
			"user_id":  member.UserID,
			"fullname": member.User.Fullname,
			"email":    member.User.Email,
			"role":     member.Role,
			"joined_at": member.CreatedAt,

		})
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

	utils.SuccessResponse(c, http.StatusOK, "Project Members Retrieved Successfully", memberResponses)
}


func (h *ProjectMemberHandler) UpdateMemberRole(c *gin.Context) {
	// HARD CODE FOR JWT
	updaterID := c.GetUint("user_id")
	
	projectIdstr := c.Param("id")
	projectID, err := strconv.ParseUint(projectIdstr, 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid project ID", err.Error())
		return
	}
	
	userIDstr := c.Param("user_id")
	userID, err := strconv.ParseUint(userIDstr, 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid user ID", err.Error())
		return
	}
	
	var req struct {
		Role string `json:"role" binding:"required"`
	}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}
	
	err = h.projectMemberService.UpdateMemberRole(uint(projectID), uint(userID), updaterID, req.Role)
	if err != nil {
		utils.ErrorResponse(c, http.StatusForbidden, "Failed to update member role", err.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Member role updated successfully", nil)
}

func (h *ProjectMemberHandler) RemoveMember(c *gin.Context) {

	removerID := c.GetUint("user_id")


	projectIdstr := c.Param("id")
	projectID, err := strconv.ParseUint(projectIdstr, 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid project ID", err.Error())
		return
	}
	userIDstr := c.Param("user_id")
	userID, err := strconv.ParseUint(userIDstr, 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid user ID", err.Error())
		return
	}
	err = h.projectMemberService.RemoveMemberFromProject(uint(projectID), uint(userID), removerID)
	if err != nil {
		
		utils.ErrorResponse(c, http.StatusForbidden, "Failed to remove member", err.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Member removed successfully", nil)
}

