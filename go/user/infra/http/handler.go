package http

import (
	"net/http"
	"sample/user/domain/context"
	"sample/user/domain/repository"
	"sample/user/iadapter/controller"
	"sample/user/iadapter/request"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func getUserList(c echo.Context) error {
	req := &request.GetUserListRequest{}

	if err := c.Bind(req); err != nil {
		return jsonResponse(c, http.StatusBadRequest, nil)
	}

	userRepo := c.Get(context.UserRepository).(repository.UserRepository)
	con := controller.NewGetUserListController(req, userRepo)

	// TODO: Error handling
	code, response, _ := con.Run()

	return jsonResponse(c, code, response)
}

func getUser(c echo.Context) error {
	req := &request.GetUserRequest{}

	if err := c.Bind(req); err != nil {
		return jsonResponse(c, http.StatusBadRequest, nil)
	}

	userID := c.Param("user_id")
	userRepo := c.Get(context.UserRepository).(repository.UserRepository)
	con := controller.NewGetUserController(req, userRepo)

	// TODO: Error handling
	code, response, _ := con.Run(userID)

	return jsonResponse(c, code, response)
}

func createUser(c echo.Context) error {
	req := &request.CreateUserRequest{}

	if err := c.Bind(req); err != nil {
		return jsonResponse(c, http.StatusBadRequest, nil)
	}

	userRepo := c.Get(context.UserRepository).(repository.UserRepository)
	vali := c.Get(context.Validator).(*validator.Validate)
	con := controller.NewCreateUserController(vali, req, userRepo)

	// TODO: Error handling
	code, response, _ := con.Run()

	return jsonResponse(c, code, response)
}

func updateUser(c echo.Context) error {
	req := &request.UpdateUserRequest{}

	if err := c.Bind(req); err != nil {
		return jsonResponse(c, http.StatusBadRequest, nil)
	}

	userID := c.Param("user_id")
	userRepo := c.Get(context.UserRepository).(repository.UserRepository)
	vali := c.Get(context.Validator).(*validator.Validate)
	con := controller.NewUpdateUserController(vali, req, userRepo)

	// TODO: Error handling
	code, response, _ := con.Run(userID)

	return jsonResponse(c, code, response)
}

func deleteUser(c echo.Context) error {
	req := &request.DeleteUserRequest{}

	if err := c.Bind(req); err != nil {
		return jsonResponse(c, http.StatusBadRequest, nil)
	}

	userID := c.Param("user_id")
	userRepo := c.Get(context.UserRepository).(repository.UserRepository)
	con := controller.NewDeleteUserController(req, userRepo)

	code, response, _ := con.Run(userID)

	return jsonResponse(c, code, response)
}
