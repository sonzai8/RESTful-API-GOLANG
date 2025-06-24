package services

import (
	"log"
	"main/internal/models"
	"main/internal/repository"
	"main/internal/utils"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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

	users, err := us.repo.FindAll(search, page, limit)

	if err != nil {
		return nil, utils.WrapError(err, "Failed to fetch users", utils.ErrCodeInternalServer)
	}
	if len(users) == 0 {
		return nil, utils.NewError("No users found", utils.ErrCodeNotFound)
	}
	var filteredUsers []models.User

	if search == "" {
		filteredUsers = users
	} else {
		search = strings.ToLower(search)
		for _, user := range users {
			name := strings.ToLower(user.Name)
			email := strings.ToLower(user.Email)
			if strings.Contains(email, search) || strings.Contains(name, search) {
				log.Print("User found: ", user.Name, " with email: ", user.Email)
				filteredUsers = append(filteredUsers, user)
			}
		}

		startIndex := (page - 1) * limit
		endIndex := startIndex + limit
		if startIndex >= len(filteredUsers) {
			return nil, utils.NewError("No users found for the given page", utils.ErrCodeNotFound)
		}
		if endIndex > len(filteredUsers) {
			endIndex = len(filteredUsers)
		}
		filteredUsers = filteredUsers[startIndex:endIndex]

	}
	return filteredUsers, nil

}
func (us *userService) CreateUser(user models.User) (models.User, error) {
	user.Email = utils.NormalizeString(user.Email)

	if _, exists := us.repo.FindByEmail(user.Email); exists {
		return models.User{}, utils.NewError("Email already exists", utils.ErrCodeEmailExists)
	}

	user.UUID = uuid.New().String()

	hashed_password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return models.User{}, utils.WrapError(err, "Failed to hash password", utils.ErrCodeInternalServer)
	}

	user.Password = string(hashed_password)

	if err := us.repo.Create(user); err != nil {
		return models.User{}, err
	}

	return user, nil
}
func (us *userService) UpdateUser(uuid string, user models.User) (models.User, error) {

	if u, exists := us.repo.FindByEmail(user.Email); exists && u.UUID != uuid {
		return models.User{}, utils.NewError("Email already exists", utils.ErrCodeEmailExists)
	}

	current_user, ok := us.repo.FindByUUID(uuid)

	if !ok {
		return models.User{}, utils.NewError("User not found", utils.ErrCodeNotFound)
	}

	current_user.Name = user.Name
	current_user.Email = utils.NormalizeString(user.Email)
	current_user.Age = user.Age
	current_user.Status = user.Status
	current_user.Level = user.Level

	if user.Password != "" {
		hashed_password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return models.User{}, utils.WrapError(err, "Failed to hash password", utils.ErrCodeInternalServer)
		}
		current_user.Password = string(hashed_password)
	}

	err := us.repo.Update(uuid, current_user)
	if err != nil {
		return models.User{}, utils.WrapError(err, "Failed to update user", utils.ErrCodeInternalServer)
	}
	return current_user, nil
}

func (us *userService) DeleteUser(uuid string) error {
	return us.repo.Delete(uuid)
}

func (us *userService) GetUserByUUID(uuid string) (models.User, error) {
	user, ok := us.repo.FindByUUID(uuid)
	if !ok {
		return models.User{}, utils.NewError("User not found", utils.ErrCodeNotFound)
	}
	return user, nil
}
