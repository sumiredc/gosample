package request

type CreateUserRequest struct {
	Name  string `json:"name" validate:"required,lte=100"`
	Email string `json:"email" validate:"required,email,lte=100"`
}
