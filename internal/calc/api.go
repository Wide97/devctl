package calc

func Calculate(op string, x, y int) (int, error) {
	switch op {
	case "add":
		return Add(x, y)
	case "sub":
		return Sub(x, y)
	case "mul":
		return Mul(x, y)
	case "div":
		return Div(x, y)
	default:
		return 0, ErrUnknowOperation
	}
}
