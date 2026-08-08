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
}
