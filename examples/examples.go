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
	fmt.Println("ST is ", ticket(config))
}

func ticket(config hash) string {
	cas := cas.NewClient(config["server"], config["username"], config["password"])
	ticket, _ := cas.RequestServiceTicket(config["service"])
	return ticket
}

func loadConfig(config string) hash {
	content := make(hash)
	raw, _ := ioutil.ReadFile(config)
	yaml.Unmarshal(raw, &content)
	return content
}
