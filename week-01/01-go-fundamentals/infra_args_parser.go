package main

import (
	"fmt"
	"os"
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
		fmt.Printf("Host: %s\n", host)
	}
}
