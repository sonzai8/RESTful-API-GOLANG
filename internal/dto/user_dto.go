package dto

import "main/internal/models"

type UserDTO struct {
	UUID   string `json:"uuid"`
	Name   string `json:"full_name" binding:"required"`
	Email  string `json:"email_address" binding:"required,email"`
	Age    int    `json:"age" binding:"required,gt=0"`
	Status string `json:"status" `
	Level  string `json:"level"`
}

type CreateUserInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email,email_advanced"`
	Age      int    `json:"age" binding:"required,gt=0"`
	Password string `json:"password" binding:"required,min=8,password_strong"`
	Status   int    `json:"status" binding:"required,oneof=1 2"`
	Level    int    `json:"level" binding:"required,oneof=1 2"`
}

type UpdateUserInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email,email_advanced"`
	Age      int    `json:"age" binding:"required,gt=0"`
	Password string `json:"password" binding:"omitempty,min=8,password_strong"`
	Status   int    `json:"status" binding:"required,oneof=1 2"`
	Level    int    `json:"level" binding:"required,oneof=1 2"`
}

func (input *CreateUserInput) MapCreateUserInputToModel() models.User {
	return models.User{

		Name:     input.Name,
		Email:    input.Email,
		Age:      input.Age,
		Password: input.Password,
		Status:   input.Status,
		Level:    input.Level,
	}
}

func (input *UpdateUserInput) MapUpdateUserInputToModel() models.User {
	return models.User{

		Name:     input.Name,
		Email:    input.Email,
		Age:      input.Age,
		Password: input.Password,
		Status:   input.Status,
		Level:    input.Level,
	}
}

func MapUserToDTO(user models.User) *UserDTO {
	return &UserDTO{
		UUID:   user.UUID,
		Name:   user.Name,
		Email:  user.Email,
		Age:    user.Age,
		Status: mapStatusText(user.Status),
		Level:  mapLevelText(user.Level),
	}
}

func MapUsersToDTOs(users []models.User) []UserDTO {
	dtos := make([]UserDTO, 0, len(users))
	for _, user := range users {

		dto := UserDTO{
			UUID:   user.UUID,
			Name:   user.Name,
			Email:  user.Email,
			Age:    user.Age,
			Status: mapStatusText(user.Status),
			Level:  mapLevelText(user.Level),
		}
		dtos = append(dtos, dto)
	}

	return dtos
}

func mapStatusText(status int) string {

	switch status {
	case 1:
		return "Show"
	case 2:
		return "Hide"
	default:
		return "None"
	}

}

func mapLevelText(level int) string {

	switch level {
	case 1:
		return "Admin"
	case 2:
		return "Member"
	default:
		return "None"
	}

}
