package http

import (
	"sample/app/user/infra/repository"
	"sample/app/user/usecase/user"

	"github.com/labstack/echo"
)

func deleteUser(e echo.Context) error {
	return nil
}

func getUserList(e echo.Context) error {
	userRepository := repository.NewUserRepository()
	usecase := user.NewGetUserListUseCase(userRepository)
	usecase.Execute()
	return nil
}

func getUser(e echo.Context) error {
	return nil
}

func saveUser(e echo.Context) error {
	return nil
}

func updateUser(e echo.Context) error {
	return nil
}
