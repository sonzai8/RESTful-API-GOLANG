package handler

import (
	"main/internal/dto"
	"main/internal/models"
	"main/internal/services"
	"main/internal/utils"
	"main/internal/validation"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service services.UserService
}

type GetUserByUUIDParams struct {
	Uuid string `uri:"uuid" binding:"required,uuid"`
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.service.GetAllUsers("", 1, 10)
	if err != nil {
		utils.ReponseErr(c, err)
		return
	}
	user_dtos := dto.MapUsersToDTOs(users)

	if len(user_dtos) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No users found"})
		return
	}

	utils.ReponseSuccess(c, http.StatusOK, user_dtos)

}

func (h *UserHandler) GetUserByUUID(c *gin.Context) {
	var params GetUserByUUIDParams
	if err := c.ShouldBindUri(&params); err != nil {
		utils.ResponseValidator(c, validation.HandleValidationErrors(err))
		return
	}
	user, err := h.service.GetUserByUUID(params.Uuid)
	if err != nil {
		utils.ReponseErr(c, err)
		return
	}
	user_dto := dto.MapUserToDTO(user)
	utils.ReponseSuccess(c, http.StatusOK, user_dto)

}

func (h *UserHandler) CreateUser(ctx *gin.Context) {
	var params models.User
	if err := ctx.ShouldBindJSON(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandleValidationErrors(err))
		return
	}

	created_user, err := h.service.CreateUser(params)

	if err != nil {
		utils.ReponseErr(ctx, err)
		return
	}
	user_dto := dto.MapUserToDTO(created_user)
	utils.ReponseSuccess(ctx, http.StatusCreated, user_dto)

}

func (h *UserHandler) UpdateUser(c *gin.Context) {

}

func (h *UserHandler) DeleteUser(c *gin.Context) {

}
