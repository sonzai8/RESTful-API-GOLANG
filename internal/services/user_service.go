package services

import (
	"main/internal/models"
	"main/internal/repository"
)

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (us *userService) GetAllUsers(search string, page, limit int) ([]models.User, error) {
	return us.repo.FindAll()
}
func (us *userService) CreateUser(user models.User) (models.User, error) {
	return us.repo.Create()
}
func (us *userService) UpdateUser(uuid string, user models.User) (models.User, error) {
	return us.repo.Update()
}
func (us *userService) DeleteUser(uuid string) error {
	return us.repo.Delete()
}
