package controller

import (
	"net/http"
	"sample/user/domain/repository"
	"sample/user/iadapter/presenter"
	"sample/user/iadapter/request"
	"sample/user/iadapter/response"
	"sample/user/usecase"
	"sample/user/usecase/dto"

	"github.com/go-playground/validator"
)

type CreateUserController struct {
	validator      *validator.Validate
	request        *request.CreateUserRequest
	userRepository repository.UserRepository
}

func NewCreateUserController(v *validator.Validate, r *request.CreateUserRequest, ur repository.UserRepository) *CreateUserController {
	return &CreateUserController{
		validator:      v,
		request:        r,
		userRepository: ur,
	}
}

func (c *CreateUserController) Run() (response.StatusCode, response.Response, error) {
	i := dto.NewCreateUserInput(c.request.Name, c.request.Email)
	u := usecase.NewCreateUserUseCase(c.userRepository)
	o, err := u.Execute(*i)

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	p := presenter.New()

	return http.StatusOK, p.User(o.User()), nil
}
