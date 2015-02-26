package service

import (
	"github.com/lucasuyezu/golang-cas-client/util"
	"strings"
)

type CasServiceConfig struct {
	Server, Service string
}

type CasServiceResponse struct {
	Status      bool
	User        string
	Email       string
	Authorities string
	CN          string
}

func New(server, service string) CasServiceConfig {
	return CasServiceConfig{server, service}
}

func (self CasServiceConfig) ValidateServiceTicket(ticket string) (*CasServiceResponse, error) {
	url := self.checkServiceTicketUrl()
	params := map[string]string{"ticket": ticket, "service": self.Service}
	body, err := util.GetResponseBody(url, params)

	if err != nil {
		return nil, err
	}

	return parseServiceResponse(body), nil
}

func (self CasServiceConfig) checkServiceTicketUrl() string {
	return self.Server + "/serviceValidate"
}

func parseServiceResponse(content string) *CasServiceResponse {
	response := CasServiceResponse{}
	response.Status = strings.Contains(content, "authenticationSuccess")
	response.User = manualExtractNode(content, "cas:user")
	response.Email = manualExtractNode(content, "cas:email")
	response.Authorities = manualExtractNode(content, "cas:authorities")
	response.CN = manualExtractNode(content, "cas:cn")

	return &response
}

func manualExtractNode(content, attr string) string {
	splitted := strings.Split(content, attr)

	if len(splitted) < 2 {
		return ""
	}

	result := strings.Replace(splitted[1], "</", "", -1)
	return strings.Replace(result, ">", "", -1)
}
