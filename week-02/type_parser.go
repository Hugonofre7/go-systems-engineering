package main

import (
	"fmt"
	"strconv"
)

func parsePort(raw string) (int, error) {
	port, err := strconv.Atoi(raw)

	if err != nil {
		return 0, fmt.Errorf("invalid port: %s", raw)
	}

	if port < 1 || port > 65535 {
		return 0, fmt.Errorf("port out of range: %s", raw)
	}

	return port, nil
}

func markUnhealthy(s Server) {
	s.Healthy = false
}

func main() {
	tests := []string{
		"8080",
		"abc",
		"99999",
	}

	for _, raw := range tests {
		port, err := parsePort(raw)

		if err != nil {
			fmt.Printf("Error parsing %q: %v\n", raw, err)
			continue
		}

		fmt.Printf("Parsed port: %d\n", port)
	}
	hostname := "café-server-01"
	fmt.Println("len:", len(hostname))

	for i, r := range hostname {
		fmt.Printf("index=%d rune=%c\n", i, r)
	}

	for i := 0; i < len(hostname); i++ {
		fmt.Printf("byte[%d]=%v\n", i, hostname[i])
	}

	var smallCounter int8 = 120
	smallCounter += 10
	fmt.Println("smallCounter after overflow:", smallCounter)

}
