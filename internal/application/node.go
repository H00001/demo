package application

import (
	"context"
	"test/internal/domain/node"
	"test/internal/infra"
)

type ServerRegisteService struct {
	*infra.Infra
	node node.NodeService
}

func NewServerRegisteService(infra *infra.Infra) *ServerRegisteService {
	svc := &ServerRegisteService{infra, node.NewNodeServer(infra)}
	return svc
}

func (svc *ServerRegisteService) Regisite(ctx context.Context, record node.ServerRecord) int {
	return svc.node.AddServerNode(ctx, record)
}

func (svc *ServerRegisteService) Cancel(ctx context.Context, serverID int) {
	svc.ServerRepo.DeleteServerNode(ctx, serverID)
}
