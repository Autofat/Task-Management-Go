package repository

import (
	"fmt"
	"task-management/internal/model"

	"gorm.io/gorm"
)

type ProjectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{db: db}
}

func (r *ProjectRepository) Create(project *model.Project) error {
	userRepo := NewUserRepository(r.db)
	_, err := userRepo.FindByID(project.OwnerID)
	if err != nil {
		return fmt.Errorf("owner with ID %d does not exist", project.OwnerID)
	}

	return r.db.Create(project).Error
}

func (r *ProjectRepository) FindByID(id uint) (*model.Project, error) {
	var project model.Project
	err := r.db.Preload("Owner").First(&project, id).Error
	if err != nil {
		return nil, err
	}
	return &project, nil
}

func (r *ProjectRepository) FindByOwnerId(ownerID uint) ([]model.Project,error) {
	var projects []model.Project
	err := r.db.Preload("Owner").Where("owner_id = ?", ownerID).Find(&projects).Error
	if err != nil {
		return nil, err
	}
	return projects, nil
}

func (r *ProjectRepository) Update(id uint, updates *model.Project) error {
	var project model.Project
	err := r.db.First(&project, id).Error
	if err != nil {
		return err
	}

	if updates.Title != "" {
		project.Title = updates.Title
	}

	if updates.OwnerID != 0 {
		userRepo := NewUserRepository(r.db)
		_,err = userRepo.FindByID(updates.OwnerID)
		if err != nil {
			return err
		}
		project.OwnerID = updates.OwnerID
	}
	return r.db.Save(&project).Error
}

func (r *ProjectRepository) DeleteById(id uint) error {
	result := r.db.Delete(&model.Project{}, id)
	taskRepo := NewTaskRepository(r.db)
	err := taskRepo.db.Where("project_id = ?", id).Delete(&model.Task{}).Error
	if err != nil {
		return err
	}
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no project found with the given ID")
	}
	return nil
}
