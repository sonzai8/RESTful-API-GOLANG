package services

import "main/internal/models"

type UserService interface {
	GetAllUsers(search string, page, limit int) ([]models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(uuid string, user models.User) (models.User, error)
	DeleteUser(uuid string) error
}
