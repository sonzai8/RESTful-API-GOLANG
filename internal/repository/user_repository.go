package repository

import (
	"main/internal/models"
)

type InMemoryUserRepository struct {
	users []models.User
}

func NewInMemoryUserRepository() UserRepository {
	return &InMemoryUserRepository{
		users: make([]models.User, 0),
	}
}

// FindAll returns all users in the repository
func (r *InMemoryUserRepository) FindAll() ([]models.User, error) {
	return r.users, nil
}

func (r *InMemoryUserRepository) FindByUUID(uuid string) (models.User, bool) {

	for _, user := range r.users {
		if user.UUID == uuid {
			return user, true
		}
	}

	return models.User{}, false
}

func (r *InMemoryUserRepository) FindByEmail(email string) (models.User, bool) {

	for _, user := range r.users {
		if user.Email == email {
			return user, true
		}
	}

	return models.User{}, false
}

func (r *InMemoryUserRepository) Create(user models.User) error {
	r.users = append(r.users, user)

	return nil
}

func (r *InMemoryUserRepository) Update() (models.User, error) {
	return models.User{}, nil
}

func (r *InMemoryUserRepository) Delete() error {
	return nil
}
