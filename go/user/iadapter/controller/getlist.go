package controller

import (
	"net/http"
	"sample/user/domain/repository"
	"sample/user/iadapter/presenter"
	"sample/user/iadapter/request"
	"sample/user/iadapter/response"
	"sample/user/usecase"
	"sample/user/usecase/dto"
)

type GetUserListController struct {
	request        *request.GetUserListRequest
	userRepository repository.UserRepository
}

func NewGetUserListController(r *request.GetUserListRequest, ur repository.UserRepository) *GetUserListController {
	return &GetUserListController{
		request:        r,
		userRepository: ur,
	}
}

func (c *GetUserListController) Run() (response.StatusCode, response.Response) {
	i := dto.NewGetUserListInput()
	u := usecase.NewGetUserListUseCase(c.userRepository)
	o, err := u.Execute(*i)

	if err != nil {
		return http.StatusInternalServerError, nil
	}

	p := presenter.New()

	return http.StatusOK, p.UserList(o.Users())
}
