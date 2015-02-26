package cas

import "github.com/lucasuyezu/golang-cas-client/client"

func NewClient(server, username, password string) client.CasClientConfig {
	return client.New(server, username, password)
}
