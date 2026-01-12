package calc

import (
	"math"
	"testing"
)

func TestAdd_OK(t *testing.T) {
	res, err := add(2, 3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res != 5 {
		t.Fatalf("expected 5, got %f", res)
	}
}

func TestADD_Overflow(t *testing.T) {
	_, err := add(math.MaxFloat64, math.MaxFloat64)
	if err != ErrAritmOverflow {
		t.Fatalf("expected ErrAritmOverflow,got %v", err)
	}
}

func TestDiv_ByZero(t *testing.T) {
	_, err := div(10, 0)
	if err != ErrDivisionByZero {
		t.Fatalf("expected ErrDivisionByZero,got %v", err)
	}
}

func TestDiv_OK(t *testing.T) {
	res, err := div(10, 2)
	if err != nil {
		t.Fatalf("unexpected error:%v", err)
	}

	if res != 5 {
		t.Fatalf("expected 5, got %f", res)
	}
}

func TestMul_Overflow(t *testing.T) {
	_, err := mul(math.MaxFloat64, math.MaxFloat64)
	if err != ErrAritmOverflow {
		t.Fatalf("expected ErrAritmOverflow,got %v", err)
	}
}

func TestMul_OK(t *testing.T) {
	res, err := mul(10, 2)
	if err != nil {
		t.Fatalf("unexpected error:%v", err)
	}

	if res != 20 {
		t.Fatalf("expected 20, got %f", res)
	}
}

func TestSub_OK(t *testing.T) {
	res, err := sub(7, 2)
	if err != nil {
		t.Fatalf("unexpected error:%v", err)
	}

	if res != 5 {
		t.Fatalf("expected 5, got %f", res)
	}
}

func TestSub_Overflow(t *testing.T) {
	_, err := sub(math.MaxFloat64, -math.MaxFloat64)
	if err != ErrAritmOverflow {
		t.Fatalf("expected ErrAritmOverflow,got %v", err)
	}
}
