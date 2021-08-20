package utils

import "fmt"

func LogOutput(message string) {
	fmt.Println(message)
}

func NewLogger() Logger {
	log := LoggerAdapter(LogOutput)
	return log
}
