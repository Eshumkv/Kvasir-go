package utils

import (
	"fmt"
)

// DEBUG prints a debug message
func DEBUG(a ...interface{}) {
	fmt.Println(a...)
}

// NOP is used to remove a warning from go.
func NOP(a ...interface{}) {
}
