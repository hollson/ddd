package aggregate

import (
	"context"

	"github.com/hollson/ddd_auth2/domain/dto"
	"github.com/hollson/ddd_auth2/domain/entity"
	"github.com/hollson/ddd_auth2/domain/repo"
	"github.com/hollson/ddd_auth2/domain/repo/specification"
	"github.com/hollson/ddd_auth2/infrastructure/pkg/hcode"
	"github.com/hollson/ddd_auth2/infrastructure/pkg/tool"
)

type AuthFactory struct {
	merchantRepo  repo.MerchantRepo
	authCodeRepo  repo.AuthCodeRepo
	authTokenRepo repo.AuthTokenRepo
}

func (f *AuthFactory) NewAuthCode(ctx context.Context, data dto.AuthCodeReq) (authCode AuthCode, err error) {
	var merchant *entity.Merchant
	var spec = specification.NewMerchantSpecificationByAPPID(data.APPID)
	if err = spec.ParameterCheck(ctx); err != nil {
		return
	}
	merchant, err = f.merchantRepo.QueryMerChant(ctx, spec)
	if err != nil {
		return
	}
	return AuthCode{
		authCodeRepo: f.authCodeRepo,
		data:         data,
		merchant:     merchant,
	}, nil
}

func (f *AuthFactory) NewProduceAuthToken(ctx context.Context, data dto.ProduceAuthTokenReq) (authToken AuthTokenProduce, err error) {
	var spec = specification.NewMerchantSpecificationByAPPID(data.APPID)
	if err = spec.ParameterCheck(ctx); err != nil {
		return
	}
	merchant, err := f.merchantRepo.QueryMerChant(ctx, spec)
	if err != nil {
		return
	}
	return AuthTokenProduce{
		authCodeRepo:  f.authCodeRepo,
		authTokenRepo: f.authTokenRepo,
		merchant:      merchant,
		data:          data,
	}, nil
}

func (f *AuthFactory) NewAuthToken(ctx context.Context, data dto.OpenIdTokenReq) (auth AuthToken, err error) {
	var (
		tokenData tool.JwtTokenData
	)
	if tokenData, err = tool.CheckAuthToken(data.Token); err != nil {
		return
	}
	if tokenData.OpenId != data.OpenId {
		return AuthToken{}, hcode.TokenValidErr
	}
	return AuthToken{
		openId:        tokenData.OpenId,
		appId:         tokenData.AppId,
		authTokenRepo: f.authTokenRepo,
	}, nil
}
