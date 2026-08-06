package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: infra_args_parser <host1> <host2> ...")
		os.Exit(1)
	}
	hosts := args[1:]
	fmt.Println("Hosts received:", len(hosts))
	validCount := 0
	invalidCount := 0
	for _, host := range hosts {
		if err := validateHost(host); err != nil {
			fmt.Printf("Error: %v\n", err)
			invalidCount++
			continue
		}
		fmt.Printf("Host: %s\n", host)
		validCount++
	}
	fmt.Printf("Valid hosts: %d\n", validCount)
	fmt.Printf("Invalid hosts: %d\n", invalidCount)
}

func validateHost(host string) error {
	if strings.TrimSpace(host) == "" {
		return fmt.Errorf("invalid host \"%s\": empty string", host)
	}
	if strings.Contains(host, " ") {
		return fmt.Errorf("invalid host \"%s\": contains spaces", host)
	}
	if !strings.Contains(host, ".") {
		return fmt.Errorf("invalid host \"%s\": missing dot", host)
	}
	return nil
}
