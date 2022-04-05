package aggregate

import (
	"context"
	"github.com/hollson/ddd_auth2/domain/dto"
	"github.com/hollson/ddd_auth2/domain/entity"
	"github.com/hollson/ddd_auth2/domain/obj"
	"github.com/hollson/ddd_auth2/domain/repo"
	"github.com/hollson/ddd_auth2/domain/repo/specification"
	consts "github.com/hollson/ddd_auth2/infrastructure/conf"
	"github.com/hollson/ddd_auth2/infrastructure/pkg/hcode"
	"github.com/hollson/ddd_auth2/infrastructure/pkg/tool"
)

type AuthTokenProduce struct {
	authCodeRepo  repo.AuthCodeRepo
	authTokenRepo repo.AuthTokenRepo
	merchant      *entity.Merchant
	data          dto.ProduceAuthTokenReq
}

func (a *AuthTokenProduce) ProduceToken(ctx context.Context) (authTokenSimple dto.AuthTokenSimple, err error) {
	var codeSpec = specification.NewAuthCodeSpecificationByCode(a.data.Code)
	if err = codeSpec.ParameterCheck(ctx); err != nil {
		return
	}
	dataCode, err := a.authCodeRepo.QueryCode(ctx, codeSpec)
	if err != nil {
		return
	}
	if dataCode.APPID != a.data.APPID {
		err = hcode.ParameterErr
		return
	}
	if a.data.Secret != a.merchant.Secret {
		err = hcode.ParameterErr
		return
	}
	var (
		data = obj.AuthToken{
			APPID:  a.data.APPID,
			Secret: a.merchant.Secret,
			OpenID: dataCode.OpenID,
			Scope:  a.merchant.Scope,
		}
		accessTokenJwt  tool.JwtToken
		refreshTokenJwt tool.JwtToken
		reqJwtToken     = tool.JwtTokenData{
			OpenId: dataCode.OpenID,
			AppId:  a.data.APPID,
			Scope:  dataCode.Scope,
		}
	)
	accessTokenJwt, err = tool.CreateAuthToken(reqJwtToken, consts.AuthAccessTokenCacheKeyTimeout)
	if err != nil {
		return
	}
	data.AccessToken = accessTokenJwt.Token
	data.AccessTokenTimeline = accessTokenJwt.TokenTimeline
	refreshTokenJwt, err = tool.CreateAuthToken(reqJwtToken, consts.AuthRefreshTokenCacheKeyTimeout)
	if err != nil {
		return
	}
	data.RefreshToken = refreshTokenJwt.Token
	data.RefreshTokenTimeline = refreshTokenJwt.TokenTimeline
	err = a.authTokenRepo.CreateAuthToken(ctx, data)
	if err != nil {
		return
	}
	_ = a.authCodeRepo.DelCode(ctx, specification.NewAuthCodeSpecificationByCode(a.data.Code))
	return data.TOSimple(), nil
}
