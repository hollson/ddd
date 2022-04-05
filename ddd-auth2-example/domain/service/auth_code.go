package service

import (
	"context"

	"github.com/hollson/ddd_auth2/domain/aggregate"
	"github.com/hollson/ddd_auth2/domain/dto"
	"github.com/hollson/ddd_auth2/domain/repo"
)

type AuthCode struct {
	factory      *aggregate.AuthFactory
	authCodeRepo repo.AuthCodeRepo
}

func (a *AuthCode) CreateCodeOpenId(ctx context.Context, req dto.AuthCodeReq) (string, error) {
	if err := req.Check(); err != nil {
		return "", err
	}
	f, err := a.factory.NewAuthCode(ctx, req)
	if err != nil {
		return "", err
	}
	return f.CreateCode(ctx)
}
