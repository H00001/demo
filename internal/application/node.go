package application

import (
	"context"
	"test/internal/domain/node"
	"test/internal/infra"
)

type ServerRegisteService struct {
	*infra.Infra
}

func NewServerRegisteService(infra *infra.Infra) *ServerRegisteService {
	svc := &ServerRegisteService{infra}
	return svc
}

func (svc *ServerRegisteService) Regisite(ctx context.Context, record node.ServerRecord) int {
	return svc.ServerRepo.AddServerNode(ctx, record)
}

func (svc *ServerRegisteService) Cancel(ctx context.Context, serverName int) {
	svc.ServerRepo.DeleteServerNode(ctx, 1)
}
