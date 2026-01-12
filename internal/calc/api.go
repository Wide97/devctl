package calc

func Calculate(op string, x, y float64) (float64, error) {
	switch op {
	case "add":
		return add(x, y)
	case "sub":
		return sub(x, y)
	case "mul":
		return mul(x, y)
	case "div":
		return div(x, y)
	default:
		return 0, ErrUnknowOperation
	}
}
