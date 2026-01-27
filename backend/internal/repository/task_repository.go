package repository

import (
	"task-management/internal/model"

	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) Create(task *model.Task) error {
	projectRepo := NewProjectRepository(r.db)
	_, err := projectRepo.FindByID(task.ProjectID)
	if err != nil{
		return err
	}
	return r.db.Create(task).Error
}

func (r *TaskRepository) FindByID(id uint, ProjectId uint) (*model.Task, error) {
	var task model.Task
	err := r.db.Preload("Project").Where("project_id = ?", ProjectId).First(&task, id).Error
	if err !=  nil {
		return nil, err
	}
	return &task, nil
}

func (r *TaskRepository) FindByProjectID(ProjectId uint) ([]model.Task, error) {
	var task []model.Task
	projectRepo := NewProjectRepository(r.db)
	_, err := projectRepo.FindByID(ProjectId)
	if err != nil{
		return nil, err
	}

	err = r.db.Preload("Project").Where("project_id = ?", ProjectId).Find(&task).Error
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (r *TaskRepository) GetTasksByProjectIDWithFilters(ProjectId uint, status, priority string, page, limit int, sort, order string) ([]model.Task, int64, error) {
	var tasks []model.Task
	var totalCount int64

	offseth := (page - 1) * limit
	query := r.db.Model(&model.Task{}).Where("project_id = ?", ProjectId)

	if status != "" {
		query = query.Where("status = ?", status)
	}
	if priority != "" {
		query = query.Where("priority = ?", priority)
	}
	if sort != "" && order != "" {
		query = query.Order(sort + " " + order)
	}else {
		query = query.Order("created_at ASC")
	}

	query.Count(&totalCount)

	err := query.Offset(offseth).Limit(limit).Preload("Project").Find(&tasks).Error
	return tasks, totalCount, err

}

func (r *TaskRepository) Update(id uint, updates *model.Task) error {
	var task model.Task
	UserRepo := NewUserRepository(r.db)
	err := r.db.First(&task, id).Error
	if err != nil {
		return err
	}
	if updates.Title != "" {
		task.Title = updates.Title
	}
	if updates.Description != "" {
		task.Description = updates.Description
	}
	if updates.Priority != "" {
		task.Priority = updates.Priority
	}
	if updates.Status != "" {
		task.Status = updates.Status
	}
	if updates.DueDate != "" {
		task.DueDate = updates.DueDate
	}
	if updates.AssignedID != 0 {
		if _, err := UserRepo.FindByID(updates.AssignedID); err != nil {
			return err
		}
		task.AssignedID = updates.AssignedID
	}

	return r.db.Save(&task).Error
}

func (r *TaskRepository) DeleteById(id uint) error {
	err := r.db.Delete(&model.Task{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
