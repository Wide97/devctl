package output

import "fmt"

func OK(msg string) {
	fmt.Printf("[OK] %s\n", msg)
}

func Error(msg string) {
	fmt.Printf("[ERROR] %s\n", msg)
}
