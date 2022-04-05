package specification

import (
	"context"

	"github.com/hollson/ddd_auth2/domain/repo"
	"github.com/hollson/ddd_auth2/infrastructure/pkg/hcode"
)

type AuthTokenByOpenId struct {
	OpenId string `json:"code"`
}

func NewAuthTokenSpecificationByoOenId(openId string) repo.AuthTokenSpecificationRepo {
	return &AuthTokenByOpenId{OpenId: openId}
}

func (m AuthTokenByOpenId) ParameterCheck(ctx context.Context) error {
	if m.OpenId == "" {
		return hcode.SysParameterErr
	}
	return nil
}

func (m AuthTokenByOpenId) ToSql(ctx context.Context) interface{} {
	return m.OpenId
}
