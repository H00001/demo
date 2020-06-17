package controllers

import "test/internal/domain/node"

type ServerNode struct {
	// server name
	ServerName string `json:"server_name"`
	// server address
	Address string `json:"address"`
	// server node provide services
	Provide []string `json:"provide"`
	// available protocols
	Protocols int `json:"protocols"`
}

func NewServerRecode(s ServerNode) node.ServerRecord {
	return node.ServerRecord{
		ServerName: s.ServerName,
		Address:    s.Address,
		Provide:    s.Provide,
		Protocols:  s.Protocols,
	}

}
