package output

import (
	"fmt"
	"io"
	"os"
)

// configurable writer -> test and future formats
var (
	stdout io.Writer = os.Stdout
	stderr io.Writer = os.Stderr
)

// PrintSuccess prints a success message from stdout
func PrintSuccess(msg string) {
	fmt.Fprintf(stdout, msg)
}

// PrintError prints an error message stderr
func PrintError(msg string) {
	fmt.Fprintf(stderr, msg)
}
