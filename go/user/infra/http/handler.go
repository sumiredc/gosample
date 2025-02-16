package http

import (
	"net/http"
	"sample/user/domain/entity"
	"sample/user/domain/repository"
	"sample/user/usecase"

	"github.com/labstack/echo"
)

type ErrorResponse struct {
	message string
}

func deleteUser(c echo.Context) error {
	return nil
}

func getUserList(c echo.Context) error {
	userRepository := c.Request().Context().Value(userRepository).(repository.UserRepository)
	usecase := usecase.NewGetUserListUseCase(userRepository)
	users, err := usecase.Execute()

	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			ErrorResponse{message: http.StatusText(http.StatusInternalServerError)},
		)
	}

	type Response struct {
		message string
		users   []entity.User
	}

	return c.JSON(http.StatusOK, Response{
		message: http.StatusText(http.StatusOK),
		users:   users,
	})
}

func getUser(c echo.Context) error {
	return nil
}

func saveUser(c echo.Context) error {
	input := &usecase.SaveUserInput{}
	if err := c.Bind(input); err != nil {
		return err
	}

	usecase := usecase.NewSaveUserUseCase(input)
	usecase.Execute()

	return nil
}

func updateUser(c echo.Context) error {
	input := &usecase.UpdateUserInput{}
	if err := c.Bind(input); err != nil {
		return err
	}

	input.UserID = c.Param("UserID")

	usecase := usecase.NewUpdateUserUseCase(input)
	usecase.Execute()

	return nil
}
