package controller

import (
	"net/http"
	"sample/user/domain/repository"
	"sample/user/iadapter/request"
	"sample/user/iadapter/response"
	"sample/user/usecase"
	"sample/user/usecase/dto"

	"github.com/go-playground/validator"
)

type UpdateUserController struct {
	validator      *validator.Validate
	request        *request.UpdateUserRequest
	userRepository repository.UserRepository
}

func NewUpdateUserController(v *validator.Validate, r *request.UpdateUserRequest, ur repository.UserRepository) *UpdateUserController {
	return &UpdateUserController{
		validator:      v,
		request:        r,
		userRepository: ur,
	}
}

func (c *UpdateUserController) Run(userId string) (response.StatusCode, response.Response) {
	if err := c.validator.Struct(c.request); err != nil {
		return http.StatusUnprocessableEntity, nil
	}

	i := dto.NewUpdateUserInput(userId, c.request.Name, c.request.Email)
	u := usecase.NewUpdateUserUseCase(c.userRepository)
	_, err := u.Execute(*i)

	if err != nil {
		return http.StatusInternalServerError, nil
	}

	return http.StatusOK, nil
}
