package main

import "fmt"

type Server struct {
	Host    string
	Port    int
	Healthy bool
}

func main() {
	registry := make(map[string]Server)

	registry["node-1"] = Server{
		Host:    "server-01.example.com",
		Port:    8080,
		Healthy: true,
	}

	registry["node-2"] = Server{
		Host:    "server-02.example.com",
		Port:    8081,
		Healthy: false,
	}

	registry["node-3"] = Server{
		Host:    "server-03.example.com",
		Port:    8082,
		Healthy: true,
	}

	for key, server := range registry {
		fmt.Printf(
			"%s → Host: %s, Port: %d, Healthy: %t\n",
			key,
			server.Host,
			server.Port,
			server.Healthy,
		)
	}
	var brokenRegistry map[string]Server

	server, exists := brokenRegistry["node-1"]
	fmt.Printf("Read from nil map -> server: %+v, exists: %t\n", server, exists)

	brokenRegistry["node-1"] = Server{
		Host:    "test",
		Port:    1234,
		Healthy: true,
	}
}
