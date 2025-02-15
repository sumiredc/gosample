package http

import (
	"sample/user/domain/repository"
	"sample/user/usecase"

	"github.com/labstack/echo"
)

func deleteUser(c echo.Context) error {
	return nil
}

func getUserList(c echo.Context) error {
	userRepository := c.Request().Context().Value("UserRepository").(repository.UserRepository)
	usecase := usecase.NewGetUserListUseCase(userRepository)
	usecase.Execute()
	return nil
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
