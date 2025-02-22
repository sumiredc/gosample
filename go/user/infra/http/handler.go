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

	ctx := c.Request().Context()
	userRepo := ctx.Value(context.UserRepository).(repository.UserRepository)
	con := controller.NewGetUserListController(req, userRepo)

	code, response := con.Run()

	return jsonResponse(c, code, response)
}

func getUser(c echo.Context) error {
	req := &request.GetUserRequest{}

	if err := c.Bind(req); err != nil {
		return jsonResponse(c, http.StatusBadRequest, nil)
	}

	userId := c.Param("UserId")
	ctx := c.Request().Context()
	userRepo := ctx.Value(context.UserRepository).(repository.UserRepository)
	con := controller.NewGetUserController(req, userRepo)

	code, response := con.Run(userId)

	return jsonResponse(c, code, response)
}

func createUser(c echo.Context) error {
	req := &request.CreateUserRequest{}

	if err := c.Bind(req); err != nil {
		return jsonResponse(c, http.StatusBadRequest, nil)
	}

	ctx := c.Request().Context()
	userRepo := ctx.Value(context.UserRepository).(repository.UserRepository)
	vali := ctx.Value(context.Validator).(*validator.Validate)
	con := controller.NewCreateUserController(vali, req, userRepo)

	code, response := con.Run()

	return jsonResponse(c, code, response)
}

func updateUser(c echo.Context) error {
	req := &request.UpdateUserRequest{}

	if err := c.Bind(req); err != nil {
		return jsonResponse(c, http.StatusBadRequest, nil)
	}

	userId := c.Param("UserId")
	ctx := c.Request().Context()
	userRepo := ctx.Value(context.UserRepository).(repository.UserRepository)
	vali := ctx.Value(context.Validator).(*validator.Validate)
	con := controller.NewUpdateUserController(vali, req, userRepo)

	code, response := con.Run(userId)

	return jsonResponse(c, code, response)
}

func deleteUser(c echo.Context) error {
	req := &request.DeleteUserRequest{}

	if err := c.Bind(req); err != nil {
		return jsonResponse(c, http.StatusBadRequest, nil)
	}

	userId := c.Param("UserId")
	ctx := c.Request().Context()
	userRepo := ctx.Value(context.UserRepository).(repository.UserRepository)
	con := controller.NewDeleteUserController(req, userRepo)

	code, response := con.Run(userId)

	return jsonResponse(c, code, response)
}
