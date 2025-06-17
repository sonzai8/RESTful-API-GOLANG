package repository

import "main/internal/models"

type UserRepository interface {
	FindAll() ([]models.User, error)
	FindByUUID(uuid string) (models.User, bool)
	FindByEmail(email string) (models.User, bool)
	Create(user models.User) error
	Update() (models.User, error)
	Delete() error
}
