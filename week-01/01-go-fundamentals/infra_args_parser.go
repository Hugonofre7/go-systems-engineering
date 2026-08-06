package main

import (
	"errors"
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
	for _, host := range hosts {
		if err := validateHost(host); err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
		fmt.Printf("Host: %s\n", host)
	}
}

func validateHost(host string) error {
	if strings.TrimSpace(host) == "" {
		return errors.New("invalid host: empty string")
	}
	if strings.Contains(host, " ") {
		return errors.New("invalid host: contains spaces")
	}
	if !strings.Contains(host, ".") {
		return errors.New("invalid host: missing dot")
	}
	return nil
}
