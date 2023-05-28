package internal

import (
	"context"
	"github.com/NetfluxESIR/backend/pkg/api/gen"
	"github.com/deepmap/oapi-codegen/pkg/types"
)

type Client struct {
	token     string
	apiClient *gen.ClientWithResponses
}

func NewClient(ctx context.Context, host string, email string, password string, role *string) *Client {
	c, err := gen.NewClientWithResponses(host)
	if err != nil {
		panic(err)
	}
	token, err := c.LoginUserWithResponse(ctx, gen.Account{
		Email:    types.Email(email),
		Password: password,
		Role:     (*gen.AccountRole)(role),
	})
	if err != nil {
		panic(err)
	}
	if token.JSON200 == nil {
		panic("token is nil")
	}
	return &Client{
		apiClient: c,
		token:     token.JSON200.Token,
	}
}
