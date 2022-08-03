package reference

import (
	"context"

	"github.com/zhaoyunxing/domain"
)

type SearchService struct {
	GetUser func(ctx context.Context, req *domain.User) (rsp *domain.User, err error)
}

type ModelService struct {
	SearchService
	//GetUser func(ctx context.Context, req *domain.User) (rsp *domain.User, err error)
}

func (model *ModelService) Reference() string {
	return "UserService"
}

func (model *SearchService) Reference() string {
	return "OrderService"
}
