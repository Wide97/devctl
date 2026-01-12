package calc

import "math"

func add(x, y float64) (float64, error) {
	res := x + y

	if math.IsInf(res, 0) || math.IsNaN(res) {
		return 0, ErrAritmOverflow
	}

	return res, nil
}

func sub(x, y float64) (float64, error) {
	res := x - y

	if math.IsInf(res, 0) || math.IsNaN(res) {
		return 0, ErrAritmOverflow
	}

	return res, nil
}

func div(x, y float64) (float64, error) {
	if y == 0 {
		return 0, ErrDivisionByZero
	}

	res := x / y

	if math.IsInf(res, 0) || math.IsNaN(res) {
		return 0, ErrInvalidOperation
	}

	return res, nil
}

func mul(x, y float64) (float64, error) {
	res := x * y

	if math.IsInf(res, 0) || math.IsNaN(res) {
		return 0, ErrAritmOverflow
	}

	return res, nil
}
