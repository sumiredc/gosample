package http

import (
	"net/http"
	"sample/user/iadapter/response"

	"github.com/labstack/echo"
)

type MessageResponse struct {
	Message string `json:"message"`
}

func NewMessageResponse(code int) MessageResponse {
	return MessageResponse{
		Message: http.StatusText(code),
	}
}

type DataResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewDataResponse(code int, data interface{}) DataResponse {
	return DataResponse{
		Message: http.StatusText(code),
		Data:    data,
	}
}

func jsonResponse(c echo.Context, code int, res response.Response) error {
	if res == nil {
		return c.JSON(code, NewMessageResponse(code))
	}

	return c.JSON(code, NewDataResponse(code, res))
}
