package calc

import "errors"

var (
	ErrDivisionByZero  = errors.New("Error division by Zero!")
	ErrUnknowOperation = errors.New("Error unknow operation!")
	ErrOverflow        = errors.New("Error overflow")
	ErrUnderflow       = errors.New("Error underflow")
)
