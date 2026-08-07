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

func main() {
	if currentLevel >= Warn {
		fmt.Println("High severity log level")
	} else {
		fmt.Println("Normal log level")
	}
}
