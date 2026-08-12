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
}
