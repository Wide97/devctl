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
	fmt.Fprintln(stdout, msg)
}

// PrintError prints an error message stderr
func PrintError(msg string) {
	fmt.Fprintln(stderr, msg)
}
