package service

import (
	"task-management/internal/model"
	"task-management/internal/repository"
)

type ProjectMemberService struct {
	userRepository *repository.UserRepository
	projectRepository *repository.ProjectRepository
	projectMemberRepository *repository.ProjectMemberRepository
}

func NewProjectMemberService(userRepo *repository.UserRepository, projectRepo *repository.ProjectRepository, projectMemberRepo *repository.ProjectMemberRepository) *ProjectMemberService {
	return &ProjectMemberService{
		userRepository: userRepo,
		projectRepository: projectRepo,
		projectMemberRepository: projectMemberRepo,
	}

}

func (s *ProjectMemberService) InviteMember(projectID , userID, inviterID uint) error {
	if projectID == 0 || userID == 0 || inviterID == 0 {
		return ErrInvalidInput
	}
	
	projects, err := s.projectRepository.FindByID(projectID)
	if err != nil {
		return ErrProjectNotFound
	}

	inviter, err := s.userRepository.FindByID(inviterID)
	if err != nil {
		return ErrUserNotFound
	}
	
	if projects.OwnerID != inviterID && inviter.Role != "ADMIN" {
		return ErrUnauthorized
	}

	_, err = s.userRepository.FindByID(userID)
	if err != nil {
		return ErrUserNotFound
	}

	isMember, err := s.projectMemberRepository.IsMember(projectID, userID)
	if err != nil {
		return err
	}
	if isMember {
		return ErrUserAlreadyExistsinProject
	}
	return s.projectMemberRepository.AddMember(projectID, userID, "MEMBER")
}

func (s *ProjectMemberService) GetProjectMembers(projectID uint) ([]model.ProjectMember, error) {
	if projectID == 0 {
		return nil, ErrInvalidInput
	}
	
	_, err := s.projectRepository.FindByID(projectID)
	if err != nil {
		return nil, ErrProjectNotFound
	}

	isMember, err := s.projectMemberRepository.IsMember(projectID, projectID)
	if err != nil {
		return nil, err
	}
	if !isMember {
		return nil, ErrUnauthorized
	}

	return s.projectMemberRepository.GetMembersByProjectID(projectID)
}

func (s *ProjectMemberService) UpdateMemberRole(projectID, userID, updaterID uint, newRole string) error {
	if projectID == 0 || userID == 0 || updaterID == 0 || newRole == "" {
		return ErrInvalidInput
	}
	updater, err := s.userRepository.FindByID(updaterID)
	if err != nil {
		return ErrUserNotFound
	}
	project, err := s.projectRepository.FindByID(projectID)
	if err != nil {
		return ErrProjectNotFound
	}
	if project.OwnerID != updaterID && updater.Role != "ADMIN" {
		return ErrUnauthorized
	}
	isMember, err := s.projectMemberRepository.IsMember(projectID, userID)
	if err != nil {
		return err
	}
	if !isMember {
		return ErrUserNotMember
	}
	return s.projectMemberRepository.UpdateMemberRole(projectID, userID, newRole)
}

func (s *ProjectMemberService) RemoveMemberFromProject(projectID, userID, removerID uint) error {
	if projectID == 0 || userID == 0 || removerID == 0 {
		return ErrInvalidInput
	}
	remover, err := s.userRepository.FindByID(removerID)
	if err != nil {
		return ErrUserNotFound
	}
	project, err := s.projectRepository.FindByID(projectID)
	if err != nil {
		return ErrProjectNotFound
	}
	if project.OwnerID != removerID && remover.Role != "ADMIN" {
		return ErrUnauthorized
	}
	
	_, err = s.userRepository.FindByID(userID)
	if err != nil {
		return ErrUserNotFound
	}

	if project.OwnerID == userID {
		return ErrCannotRemoveOwner
	}

	isMember, err := s.projectMemberRepository.IsMember(projectID, userID)
	if err != nil {
		return err
	}
	if !isMember {
		return ErrUserNotMember
	}
	
	return s.projectMemberRepository.RemoveMember(projectID, userID)
}

func (s *ProjectMemberService) IsMember(projectID, userID uint) (bool, error) {
	if projectID == 0 || userID == 0 {
		return false, ErrInvalidInput
	}
	return s.projectMemberRepository.IsMember(projectID, userID)
}


