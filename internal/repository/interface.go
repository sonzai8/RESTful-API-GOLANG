package repository

import "main/internal/models"

type UserRepository interface {
	FindAll() ([]models.User, error)
	findByUUID() (models.User, bool)
	Create() (models.User, error)
	Update() (models.User, error)
	Delete() error
}
