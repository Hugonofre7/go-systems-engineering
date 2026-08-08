package main

import "fmt"

type LogLevel int

const (
	Debug LogLevel = iota
	Info
	Warn
	Error
)

var currentLevel LogLevel = Info

type ServiceFlag int

const (
	FlagTLS ServiceFlag = 1 << iota
	FlagAuth
	FlagRateLimited
)

var activeFlags ServiceFlag = FlagTLS | FlagAuth

func main() {
	if currentLevel >= Warn {
		fmt.Println("High severity log level")
	} else {
		fmt.Println("Normal log level")
	}

	if activeFlags&FlagAuth != 0 {
		fmt.Println("Authentication is enabled")
	}
	var err error = nil

	if true {
		value, err := someFunction()
		fmt.Printf("Value inside if: %s\n", value)
		fmt.Printf("Error inside if: %v\n", err)
	}

	fmt.Printf("Error after if: %v\n", err)
}

func someFunction() (string, error) {
	return "failed", fmt.Errorf("validation error")
}
