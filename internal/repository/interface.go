package repository

import "main/internal/models"

type UserRepository interface {
	FindAll(search string, page, limit int) ([]models.User, error)
	FindByUUID(uuid string) (models.User, bool)
	FindByEmail(email string) (models.User, bool)
	Create(user models.User) error
	Update(uuid string, user models.User) error
	Delete(uuid string) error
}
