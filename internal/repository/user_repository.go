package repository

import (
	"fmt"
	"main/internal/models"
	"slices"
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
func (r *InMemoryUserRepository) FindAll(search string, page, limit int) ([]models.User, error) {

	return r.users, nil
}

func (r *InMemoryUserRepository) FindByUUID(uuid string, user *models.User) error {

	for _, user := range r.users {
		if user.UUID == uuid {
			return nil
		}
	}

	return nil
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

func (r *InMemoryUserRepository) Update(uuid string, current_user models.User) error {
	for i, user := range r.users {
		if user.UUID == uuid {
			fmt.Printf("Updating user with UUID: %s\n", uuid)
			fmt.Printf("Current user data: %+v\n", current_user)
			fmt.Printf("Existing user data: %+v\n", user)
			r.users[i] = current_user // Update the user with the new data
			return nil
		}
	}
	return fmt.Errorf("user with UUID %s not found", uuid)
}

func (r *InMemoryUserRepository) Delete(uuid string) error {
	for i, user := range r.users {
		if user.UUID == uuid {
			// r.users = append(r.users[:i], r.users[i+1:]...) // Remove the user from the slice
			r.users = slices.Delete(r.users, i, i+1) // Remove the user from the slice new version
			return nil
		}
	}
	return fmt.Errorf("user with UUID %s not found", uuid)
}
