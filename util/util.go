package util

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func GetResponseHeader(url, key string, params map[string]string) (string, error) {
	client := httpClient()
	response, err := client.PostForm(url, mapToUrlValues(params))

	if err != nil {
		return "", err
	}

	header := response.Header.Get("Location")

	if header == "" {
		return "", errors.New("Header: " + key + " not found")
	}

	return header, nil
}

func GetResponseBody(url string, params map[string]string) (string, error) {
	client := httpClient()
	response, err := client.PostForm(url, mapToUrlValues(params))
	if err != nil {
		return "", err
	}

	if response.StatusCode != 200 {
		errMsg := fmt.Sprintf("response should be 200 but is: %d", response.StatusCode)
		return "", errors.New(errMsg)
	}

	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	return string(body), nil
}

func httpClient() *http.Client {
	transport := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	return &http.Client{Transport: transport}
}

func mapToUrlValues(hash map[string]string) url.Values {
	values := url.Values{}

	for key, value := range hash {
		values.Add(key, value)
	}

	return values
}
