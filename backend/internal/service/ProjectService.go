package service

import (
	"task-management/internal/model"
	"task-management/internal/repository"
)

type ProjectService struct {
	userRepository *repository.UserRepository
	projectRepository *repository.ProjectRepository
}

func NewProjectService(userRepo *repository.UserRepository, projectRepo *repository.ProjectRepository) *ProjectService {
	return &ProjectService{
		userRepository: userRepo,
		projectRepository: projectRepo,
	}
}

func (s *ProjectService) CreateProject (title string, ownerID uint) (*model.Project, error) {
	if	title == ""{
		return nil, ErrInvalidInput
	}

	_, err := s.userRepository.FindByID(ownerID)
	if err != nil {
		return nil, ErrUserNotFound
	}

	project := &model.Project{
		Title:       title,
		OwnerID:     ownerID,
	}

	err = s.projectRepository.Create(project)
	if err != nil {
		return nil, err
	}

	return project, nil
}

func (s *ProjectService) GetProjectByID(id uint) (*model.Project, error) {
	return s.projectRepository.FindByID(id)
}

func (s *ProjectService) GetProjectsByOwnerID(ownerID uint) ([]model.Project, error) {
	return s.projectRepository.FindByOwnerId(ownerID)
}

func (s *ProjectService) UpdateProject(id uint, title string, ownerID uint) error {
	if title == "" {
		return ErrInvalidInput
	}

	_, err := s.projectRepository.FindByID(id)
	if err != nil {
		return ErrProjectNotFound
	}

	_, err = s.userRepository.FindByID(ownerID)
	if err != nil {
		return ErrUserNotFound
	}

	updates := &model.Project{
		Title:   title,
	}

	return s.projectRepository.Update(id, updates)
}

func (s *ProjectService) DeleteProject(id uint) error {
	_, err := s.projectRepository.FindByID(id)
	if err != nil {
		return ErrProjectNotFound
	}
	return s.projectRepository.DeleteByID(id)
}