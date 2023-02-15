package logger

import (
	"fmt"
)

func LogInfo(module string, msg string) {
	fmt.Println(msg)
}

func LogError(module string, msg string) {
	fmt.Println(msg)
}

func LogWarning(module string, msg string) {
	fmt.Println(msg)
}
