package repository

import (
	"task-management/internal/model"

	"gorm.io/gorm"
)

type ProjectMemberRepository struct {
	db *gorm.DB
}

func NewProjectMemberRepository(db *gorm.DB) *ProjectMemberRepository {
	return &ProjectMemberRepository{db: db}
}

func (r *ProjectMemberRepository) AddMember(projectID uint, userID uint, role string) error{
	member := &model.ProjectMember{
		ProjectID: projectID,
		UserID:    userID,
		Role:      role,
	}
	return r.db.Create(member).Error
}

func (r *ProjectMemberRepository) GetMembersByProjectID(projectID uint) ([]model.ProjectMember, error){
	var members []model.ProjectMember
	err := r.db.Preload("User").Where("project_id = ?", projectID).Find(&members).Error
	if err != nil {
		return nil, err
	}
	return members, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*model.User, error) {
    var user model.User
    err := r.db.Where("email = ? AND deleted_at IS NULL", email).First(&user).Error
    return &user, err
}

func (r *ProjectMemberRepository) RemoveMember(projectID, userID uint) error{
	return r.db.Where("project_id = ? AND user_id = ?", projectID, userID).Delete(&model.ProjectMember{}).Error
}

func (r *ProjectMemberRepository) UpdateMemberRole(projectID, userID uint, updateRole string) error{
	var member model.ProjectMember
	err := r.db.Where("project_id = ? AND user_id = ?", projectID, userID).First(&member).Error
	if err != nil {
		return err
	}
	member.Role = updateRole
	return r.db.Save(&member).Error
}

func (r *ProjectMemberRepository) IsMember(projectID, userID uint) (bool, error){
	var count int64
	err := r.db.Model(&model.ProjectMember{}).Where("project_id = ? AND user_id = ?", projectID, userID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}