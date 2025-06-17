package repository

import "main/internal/models"

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
	return nil, nil
}

func (r *InMemoryUserRepository) findByUUID() (models.User, bool) {
	return models.User{}, false
}

func (r *InMemoryUserRepository) Create() (models.User, error) {
	return models.User{}, nil
}

func (r *InMemoryUserRepository) Update() (models.User, error) {
	return models.User{}, nil
}

func (r *InMemoryUserRepository) Delete() error {
	return nil
}
