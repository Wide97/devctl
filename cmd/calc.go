package cmd

import (
	"devctl/internal/calc"
	"errors"
	"fmt"
	"os"
)

func Calc() error {
	if len(os.Args) < 5 {
		return errors.New("usage: devctl calc op val1 val2")
	}

	if os.Args[1] != "calc" {
		return errors.New("unknown command")
	}

	res, err := calc.Calculate(os.Args[2], os.Args[3], os.Args[4])
	if err != nil {
		return err
	}

	fmt.Printf("The result is:%d\n", res)
	return nil
}
