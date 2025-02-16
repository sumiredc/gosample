package http

import "sample/user/domain/entity"

type errorResponse struct {
	Message string `json:"message"`
}

type getUserListResponse struct {
	Message string        `json:"message"`
	Users   []entity.User `json:"Users"`
}
