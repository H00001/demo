package application

import (
	"context"
	"test/internal/domain/user/entity"
	"test/internal/infra"
)

type UserService struct {
	*infra.Infra
}

func NewUserSvc(infraer *infra.Infra) *UserService {
	svc := &UserService{infraer}
	return svc
}

func (svc *UserService) GetUserInfo(ctx context.Context, username string) (user entity.User) {
	user = svc.UserRepo.FindByUserName(ctx, username)

	return
}
