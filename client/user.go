package client

import (
	"context"
	"fmt"

	"github.com/zhaoyunxing/domain"
	"github.com/zhaoyunxing/reference"
)

type ServiceClient struct {
	UserService *reference.UserService
}

func (s *ServiceClient) SetService(service *reference.UserService) {
	s.UserService = service
}

func (s *ServiceClient) Get(ctx context.Context, req *domain.User) (rsp *domain.User, err error) {
	fmt.Printf("ServiceClient=%v", s)
	return s.UserService.GetUser(ctx, req)
}
