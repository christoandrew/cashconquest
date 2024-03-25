package requests

import (
	"github.com/christo-andrew/maybe-go/internal/models"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"userName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (r *CreateUserRequest) User() *models.User {
	return &models.User{
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Username:  r.Username,
		Email:     r.Email,
		Password:  r.Password,
	}
}
