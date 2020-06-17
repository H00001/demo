package application

import (
	"context"
	"test/internal/domain/user/entity"
	"test/internal/infra"
)

type ServerRegisteService struct {
	*infra.Infra
}

func NewServerRegisteService(infraer *infra.Infra) *ServerRegisteService {
	svc := &ServerRegisteService{infraer}
	return svc
}

func (svc *ServerRegisteService) Regisite(ctx context.Context, record entity.ServerRecord) int {
	return svc.ServerRepo.AddServerNode(ctx, record)
}

func (svc *ServerRegisteService) Cancel(ctx context.Context, serverName int) {
	svc.ServerRepo.DeleteServerNode(ctx, 1)
}
