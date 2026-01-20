package service

import (
	"task-management/internal/model"
	"task-management/internal/repository"
)

type TaskService struct {
	taskRepository *repository.TaskRepository
	projectRepository *repository.ProjectRepository
	userRepository *repository.UserRepository
}

func NewTaskService(taskRepo *repository.TaskRepository, projectRepo *repository.ProjectRepository, userRepo *repository.UserRepository) *TaskService {
	return &TaskService{
		taskRepository: taskRepo,
		projectRepository: projectRepo,
		userRepository: userRepo,
	}
}

func (s *TaskService) CreateTask(title, description, priority string, projectID, assigneeID uint) (*model.Task, error) {
	if title == "" || description == ""  {
		return nil, ErrInvalidInput
	}
	_, err := s.projectRepository.FindByID(projectID)
	if err != nil {
		return nil, ErrProjectNotFound
	}

	if assigneeID !=0 {
		_, err = s.userRepository.FindByID(assigneeID)
		if err != nil {
			return nil, ErrUserNotFound
		}
	}

	if priority == "" {
		priority = "Medium"
	}

	status := "To Do"

	task := &model.Task{
		Title:       title,
		Description: description,
		ProjectID:   projectID,
		AssignedID:  assigneeID,
		Status:      status,
		Priority:    priority,
	}
	err = s.taskRepository.Create(task)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (s *TaskService) GetTaskByID(id, projectID uint) (*model.Task, error) {
	return s.taskRepository.FindByID(id, projectID)
}

func (s *TaskService) GetTasksByProjectID(projectID uint) ([]model.Task, error) {
	return s.taskRepository.FindByProjectID(projectID)
}

func (s *TaskService) UpdateTask(id uint, title, description, status, dueDate, priority string, projectID, assigneeID uint, ) error {
	_, err := s.taskRepository.FindByID(id, projectID)
	if err != nil {
		return ErrTaskNotFound
	}
	
	if assigneeID != 0 {
		_, err = s.userRepository.FindByID(assigneeID)
		if err != nil {
			return ErrUserNotFound
		}
	}

	updates := &model.Task{}

	if title != ""{
		updates.Title = title
	}
	if description != ""{
		updates.Description = description
	}
	if dueDate != ""{
		updates.DueDate = dueDate
	}
	if priority != ""{
		updates.Priority = priority
	}
	if status != "" {
		updates.Status = status
	}
	if assigneeID != 0{
		updates.AssignedID = assigneeID
	}

	return s.taskRepository.Update(id, updates)
}

func (s *TaskService) DeleteTask(id, projectID uint) error {
	_, err := s.taskRepository.FindByID(id, projectID)
	if err != nil {
		return ErrTaskNotFound
	}
	return s.taskRepository.DeleteById(id)
}