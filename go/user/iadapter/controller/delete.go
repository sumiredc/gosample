package controller

import (
	"errors"
	"net/http"
	"sample/user/domain/errkit"
	"sample/user/domain/repository"
	"sample/user/iadapter/request"
	"sample/user/iadapter/response"
	"sample/user/usecase"
	"sample/user/usecase/dto"
)

type DeleteUserController struct {
	request        *request.DeleteUserRequest
	userRepository repository.UserRepository
}

func NewDeleteUserController(r *request.DeleteUserRequest, ur repository.UserRepository) *DeleteUserController {
	return &DeleteUserController{
		request:        r,
		userRepository: ur,
	}
}

func (c *DeleteUserController) Run(userID string) (response.StatusCode, response.Response, error) {
	i := dto.NewDeleteUserInput(userID)
	u := usecase.NewDeleteUserUseCase(c.userRepository)
	_, err := u.Execute(*i)

	if errors.Is(err, errkit.ErrNotFound) {
		return http.StatusNotFound, nil, err
	}

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, nil, nil
}
