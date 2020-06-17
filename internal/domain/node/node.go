package node

import (
	"context"
	"test/internal/infra"
)

type ServerRecord struct {
	// server name
	ServerName string `json:"server_name"`
	// server address
	Address string `json:"address"`
	// server node provide services
	Provide []string `json:"address"`
	// available protocols
	Protocols int `json:"protocols"`
}

type NodeService interface {
	AddServerNode(context.Context, ServerRecord) int
	DeleteServerNode(context.Context, int)
}

func NewNodeServer(infra *infra.Infra) NodeService {
	v := nodeServer{i: infra}
	return &v
}

type nodeServer struct {
	i *infra.Infra
}

func (n *nodeServer) AddServerNode(context context.Context, s ServerRecord) int {
	n.i.ServerRepo.AddServerNode(context, s)
}

func (n *nodeServer) DeleteServerNode(ctx context.Context, id int) {
	n.i.ServerRepo.DeleteServerNode(ctx, id)
}
