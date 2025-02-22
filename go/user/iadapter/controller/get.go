package controller

import (
	"errors"
	"net/http"
	"sample/user/domain/errkit"
	"sample/user/domain/repository"
	"sample/user/iadapter/presenter"
	"sample/user/iadapter/request"
	"sample/user/iadapter/response"
	"sample/user/usecase"
	"sample/user/usecase/dto"
)

type GetUserController struct {
	request        *request.GetUserRequest
	userRepository repository.UserRepository
}

func NewGetUserController(r *request.GetUserRequest, ur repository.UserRepository) *GetUserController {
	return &GetUserController{
		request:        r,
		userRepository: ur,
	}
}

func (c *GetUserController) Run(userID string) (response.StatusCode, response.Response, error) {
	i := dto.NewGetUserInput(userID)
	u := usecase.NewGetUserUseCase(c.userRepository)
	o, err := u.Execute(*i)

	if errors.Is(err, errkit.ErrNotFound) {
		return http.StatusNotFound, nil, err
	}

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	p := presenter.New()

	return http.StatusOK, p.User(o.User()), nil
}
