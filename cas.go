package cas

import (
	"github.com/lucasuyezu/golang-cas-client/client"
	"github.com/lucasuyezu/golang-cas-client/service"
)

func NewClient(server, username, password string) client.CasClientConfig {
	return client.New(server, username, password)
}

func NewService(server, hostService string) service.CasServiceConfig {
	return service.New(server, hostService)
}
