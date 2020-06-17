package node

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
