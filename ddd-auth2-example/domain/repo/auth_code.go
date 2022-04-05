package repo

import (
	"context"
	"github.com/hollson/ddd_auth2/domain/obj"
)

type AuthCodeRepo interface {
	CreateCode(ctx context.Context, data obj.CodeOpenId) error
	DelCode(ctx context.Context, repo AuthCodeSpecificationRepo) error
	QueryCode(ctx context.Context, repo AuthCodeSpecificationRepo) (data obj.CodeOpenId, err error)
}

type AuthCodeSpecificationRepo interface {
	ParameterCheck(ctx context.Context) error
	ToSql(ctx context.Context) interface{}
}
