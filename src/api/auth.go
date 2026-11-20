package api

import (
	"context"
	"github.com/danielgtaylor/huma/v2"
	"log"
	"net/http"
)

var createUserOp = huma.Operation{
	OperationID: "auth-register",
	Method:      http.MethodPost,
	Path:        "/register",
	Summary:     "Register User",
	Tags:        []string{"Auth"},
}

type CreateUserInput struct {
	Body struct {
		Email    string `json:"email" example:"mail@example.com"`
		Name     string `json:"name" example:"John Janes"`
		Password string `json:"password" example:"a1s2d3f4"`
	}
}

type CreateUserOutput struct{}

func (a App) CreateUserHandler(ctx context.Context, input *CreateUserInput) (*CreateUserOutput, error) {
	err := a.Auth.CreateUserController(ctx, input.Body.Email, input.Body.Name, input.Body.Password)
	if err != nil {
		log.Println("%s", err)
	}

	return &CreateUserOutput{}, nil
}
