package calc

import "errors"

var (
	ErrDivisionByZero   = errors.New("Error division by Zero")
	ErrUnknowOperation  = errors.New("Error unknow operation")
	ErrAritmOverflow    = errors.New("Error number overflow")
	ErrInvalidOperation = errors.New("Error invalid result")
)
