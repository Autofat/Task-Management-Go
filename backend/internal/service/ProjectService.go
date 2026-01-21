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
	if ownerID == 0 {
		return nil, ErrInvalidInput
	}
	if _, err := s.userRepository.FindByID(ownerID); err != nil {
		return nil, ErrUserNotFound
	}
	return s.projectRepository.FindByOwnerId(ownerID)
}

func (s *ProjectService) UpdateProject(id uint, title string) error {
	updates, err := s.projectRepository.FindByID(id)
	if err != nil {
		return ErrProjectNotFound
	}

	if title == "" {
		return ErrInvalidInput
	}else{
		updates.Title = title
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