package main

import "fmt"

type Server struct {
	Host    string
	Port    int
	Healthy bool
}

func markUnhealthy(s Server) {
	s.Healthy = false
}

func markUnhealthyPtr(s *Server) {
	s.Healthy = false
}

func main() {
	registry := make(map[string]Server)

	servers := []Server{
		{
			Host:    "server-01.example.com",
			Port:    8080,
			Healthy: true,
		},
		{
			Host:    "server-02.example.com",
			Port:    8081,
			Healthy: false,
		},
		{
			Host:    "server-03.example.com",
			Port:    8082,
			Healthy: true,
		},
	}

	unhealthyHosts := healthCheck(servers)

	fmt.Println("Unhealthy hosts:", unhealthyHosts)

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

	markUnhealthy(registry["node-1"])
	fmt.Printf("After markUnhealthy: %+v\n", registry["node-1"])

	server := registry["node-1"]
	markUnhealthyPtr(&server)

	registry["node-1"] = server

	fmt.Printf("After markUnhealthyPtr: %+v\n", registry["node-1"])

	var brokenRegistry map[string]Server

	server, exists := brokenRegistry["node-1"]
	fmt.Printf("Read from nil map -> server: %+v, exists: %t\n", server, exists)

	// Reto 2 — INTENCIONAL: escribir en un nil map provoca panic.
	brokenRegistry["node-1"] = Server{
		Host:    "test",
		Port:    1234,
		Healthy: true,
	}
}

func healthCheck(servers []Server) []string {
	var unhealthy []string

	for _, s := range servers {
		if !s.Healthy {
			unhealthy = append(unhealthy, s.Host)
		}
	}
	return unhealthy
}
