package calc

import "errors"

var (
	errDivisionByZero  = errors.New("Error division by Zero!")
	errUnknowOperation = errors.New("Error unknow operation!")
	errOverflow        = errors.New("Error overflow")
	errUnderflow       = errors.New("Error underflow")
)
