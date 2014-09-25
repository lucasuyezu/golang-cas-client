package cas

import (
	"crypto/tls"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type CasConfig struct {
	Server, Username, Password string
}

func New(server, username, password string) CasConfig {
	return CasConfig{server, username, password}
}

func (self CasConfig) RequestServiceTicket(service string) (string, error) {
	tgt, err := self.requestTgtLocation()
	if err != nil {
		return "", err
	}

	return self.getServiceTicket(tgt, service)
}

func (self CasConfig) requestTgtLocation() (string, error) {
	params := url.Values{"username": {self.Username}, "password": {self.Password}}
	response, err := postFormData(self.Server+"/v1/tickets", params)
	defer response.Body.Close()
	if err != nil {
		return "", err
	}

	return response.Header.Get("Location"), nil
}

func (self CasConfig) getServiceTicket(tgt, service string) (string, error) {
	response, err := postFormData(tgt, url.Values{"service": {service}})
	if err != nil {
		return "", err
	}

	serviceTicket, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	if response.StatusCode != 200 {
		return "", errors.New("ST response should be 200 but is: " + string(response.StatusCode))
	}

	return string(serviceTicket), nil
}

func postFormData(url string, formData url.Values) (*http.Response, error) {
	transport := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	client := &http.Client{Transport: transport}

	response, err := client.PostForm(url, formData)
	if err != nil {
		return nil, err
	}

	return response, nil
}
