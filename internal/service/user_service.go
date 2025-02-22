package service

import (
	"github.com/nicolas-calvario/Go-Api-Crud/internal/domain"
	"github.com/nicolas-calvario/Go-Api-Crud/internal/repository"
)

type UserService interface {
	Create(user *domain.User) error
	GetByID(id int64) (*domain.User, error)
	GetAll() ([]domain.User, error)
	Update(user *domain.User) error
	Delete(id int64) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) Create(user *domain.User) error {
	return s.repo.Create(user)
}

func (s *userService) GetByID(id int64) (*domain.User, error) {
	return s.repo.GetByID(id)
}

func (s *userService) GetAll() ([]domain.User, error) {
	return s.repo.GetAll()
}

func (s *userService) Update(user *domain.User) error {
	return s.repo.Update(user)
}

func (s *userService) Delete(id int64) error {
	return s.repo.Delete(id)
}
