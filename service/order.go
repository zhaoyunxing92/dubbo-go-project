package service

import (
	"context"
	"time"

	"github.com/zhaoyunxing/domain"
)

type OrderService struct {
}

func (u *OrderService) GetOrder(ctx context.Context, req *domain.User) (*domain.User, error) {
	return &domain.User{ID: "A001", Name: req.Name, Age: 18, Time: time.Now()}, nil
}

type UserService struct {
}

func (u *UserService) GetUser(ctx context.Context, req *domain.User) (*domain.User, error) {
	return &domain.User{ID: "A001", Name: req.Name, Age: 18, Time: time.Now()}, nil
}
