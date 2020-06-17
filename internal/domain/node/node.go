package node

type Protocol int

type ServerRecord struct {
	// server name
	ServerName string `json:"server_name"`
	// server address
	Address string `json:"address"`
	// server node provide services
	Provide []string `json:"address"`
	// available protocols
	Protocols []Protocol `json:"protocols"`
}
