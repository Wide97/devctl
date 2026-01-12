package cmd

import (
	"devctl/internal/calc"
	"devctl/pkg/output"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var numberRegex = regexp.MustCompile(`^-?(0|[1-9]\d*)(\.\d+)?$`)

func Calc() error {
	if len(os.Args) < 5 {
		return errors.New("usage: devctl calc <add|sub|mul|div> <val1> <val2>")
	}

	if os.Args[1] != "calc" {
		return errors.New("unknown command")
	}

	// validate numeric format
	if !numberRegex.MatchString(os.Args[3]) {
		return fmt.Errorf("invalid number format: %s", os.Args[3])
	}
	if !numberRegex.MatchString(os.Args[4]) {
		return fmt.Errorf("invalid number format: %s", os.Args[4])
	}

	// parse strings -> float64 (ONLY correct way)
	x, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		return fmt.Errorf("invalid number: %s", os.Args[3])
	}

	y, err := strconv.ParseFloat(os.Args[4], 64)
	if err != nil {
		return fmt.Errorf("invalid number: %s", os.Args[4])
	}

	res, err := calc.Calculate(os.Args[2], x, y)
	if err != nil {
		return err
	}

	output.PrintSuccess("The result is: " + fmt.Sprintf("%.2f", res))
	return nil
}
