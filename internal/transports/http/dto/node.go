package dto

import "test/internal/domain/user/entity"

// in fact we can use domain/entity/node/ServerNode to instead it
type ServerNode struct {
	// server name
	ServerName string `json:"server_name"`
	// server address
	Address string `json:"address"`
	// server node provide services
	Provide []string `json:"provide"`
	// available protocols
	Protocols []entity.Protocol `json:"protocols"`
}

func NewServerRecode(s ServerNode) entity.ServerRecord {
	return entity.ServerRecord{
		ServerName: s.ServerName,
		Address:    s.Address,
		Provide:    s.Provide,
		Protocols:  s.Protocols,
	}

}
