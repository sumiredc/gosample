package controller

import (
	"net/http"
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

func (c *DeleteUserController) Run(userId string) (response.StatusCode, response.Response) {
	i := dto.NewDeleteUserInput(userId)
	u := usecase.NewDeleteUserUseCase(c.userRepository)
	_, err := u.Execute(*i)

	if err != nil {
		return http.StatusInternalServerError, nil
	}

	return http.StatusOK, nil
}
