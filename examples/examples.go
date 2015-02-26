package main

import (
	"fmt"
	"github.com/lucasuyezu/golang-cas-client"
	"gopkg.in/yaml.v1"
	"io/ioutil"
)

type hash map[string]string

func main() {
	config := loadConfig("examples/examples.yml")
	ticket := ticket(config)
	fmt.Println("ST is ", ticket)

	fmt.Println("valid ticket ", validate(config, ticket))
	fmt.Println("invalid ticket ", validate(config, ticket+"kdoaskd"))
}

func ticket(config hash) string {
	cas := cas.NewClient(config["server"], config["username"], config["password"])
	ticket, _ := cas.RequestServiceTicket(config["service"])
	return ticket
}

func validate(config hash, ticket string) bool {
	cas := cas.NewService(config["server"], config["service"])
	result, err := cas.ValidateServiceTicket(ticket)

	if err != nil {
		fmt.Println(err)
		return false
	}

	fmt.Println(result)

	return result.Status
}

func loadConfig(config string) hash {
	content := make(hash)
	raw, _ := ioutil.ReadFile(config)
	yaml.Unmarshal(raw, &content)
	return content
}
