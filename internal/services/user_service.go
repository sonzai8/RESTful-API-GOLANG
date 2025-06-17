package services

import (
	"main/internal/models"
	"main/internal/repository"
	"main/internal/utils"

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

	return us.repo.FindAll()
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
	return us.repo.Update()
}
func (us *userService) DeleteUser(uuid string) error {
	return us.repo.Delete()
}
func (us *userService) GetUserByUUID(uuid string) (models.User, error) {
	user, ok := us.repo.FindByUUID(uuid)
	if !ok {
		return models.User{}, utils.NewError("User not found", utils.ErrCodeNotFound)
	}
	return user, nil
}
