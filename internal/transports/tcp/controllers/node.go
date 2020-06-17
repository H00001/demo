package controllers

import (
	"context"
	"encoding/json"
	"test/internal/application"
	"test/internal/infra"
)

type TcpNodeController struct {
	infra           *infra.Infra
	registerService *application.ServerRegisteService
}

func (t *TcpNodeController) Regisite(b []byte) {
	n := ServerNode{}
	json.Unmarshal(b, &n)
	t.registerService.Regisite(context.Background(), NewServerRecode(n))
}
