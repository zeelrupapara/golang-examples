package main

import (
	"service_discovery/discovery"

	consul "github.com/hashicorp/consul/api"
)

func main() {
	config := consul.DefaultConfig()
	config.Address = "localhost:8500"

	cli, err := discovery.NewClient(config, "discovery", "localhost", 8080)
	if err != nil {
		panic(err)
	}
	if err := discovery.Exec(cli); err != nil {
		panic(err)
	}
}
