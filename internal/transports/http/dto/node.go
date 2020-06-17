package dto

import (
	"test/internal/domain/node"
)

// in fact we can use domain/entity/node/ServerNode to instead it
type ServerNode struct {
	// server name
	ServerName string `json:"server_name"`
	// server address
	Address string `json:"address"`
	// server node provide services
	Provide []string `json:"provide"`
	// available protocols
	Protocols []node.Protocol `json:"protocols"`
}

func NewServerRecode(s ServerNode) node.ServerRecord {
	return node.ServerRecord{
		ServerName: s.ServerName,
		Address:    s.Address,
		Provide:    s.Provide,
		Protocols:  s.Protocols,
	}

}
