package cmd

import (
	"os"
	"testing"
)

func TestCmdCalcOpsOk(t *testing.T) {
	oldArgs := os.Args

	defer func() { os.Args = oldArgs }()

	os.Args = []string{"devctl", "calc", "add", "2", "3"}

	if err := Calc(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	os.Args = []string{"devctl", "calc", "sub", "-2", "-5"}

	if err := Calc(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	os.Args = []string{"devctl", "calc", "div", "3", "-2"}

	if err := Calc(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	os.Args = []string{"devctl", "calc", "mul", "-1", "5"}

	if err := Calc(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestCorrectArgsNumber(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"devctl", "cal", "c", "add", "1", "2"}

	if err := Calc(); err == nil {
		t.Fatalf("expected usage error, less args required")
	}

	os.Args = []string{"devctl", "1", "2"}

	if err := Calc(); err == nil {
		t.Fatalf("expected usage error, more args required")
	}

	os.Args = []string{"devctl", "calc", "add", "1", "2"}

	if err := Calc(); err != nil {
		t.Fatalf("unexpected error, right args number, got %v", err)
	}
}

func TestCmdArgsFormat(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"devctl", "pippo", "add", "1", "2"}

	if err := Calc(); err == nil {
		t.Fatalf("expected calc unknown command error")
	}

	os.Args = []string{"devctl", "calc", "prop", "1", "2"}

	if err := Calc(); err == nil {
		t.Fatalf("expected unknown ops command error")
	}

	os.Args = []string{"devctl", "calc", "add", "f", "1"}

	if err := Calc(); err == nil {
		t.Fatalf("expected wrong float number format error")
	}

	os.Args = []string{"devctl", "calc", "add", "ciao", "1"}

	if err := Calc(); err == nil {
		t.Fatalf("expected wrong float number format error")
	}

	os.Args = []string{"devctl", "calc", "add", "c12iao", "1"}

	if err := Calc(); err == nil {
		t.Fatalf("expected wrong float number format error")
	}
}
