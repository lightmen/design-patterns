package strategy

import "testing"

func TestStrategy(t *testing.T) {
	add := new(Add)
	sub := new(Subtrack)
	mul := new(Multiply)

	calc := new(Calculator)

	x := 4
	y := 2
	calc.SetStrategy(add)
	result := calc.Calculate(x, y)
	t.Logf("add result: %d\n", result)

	calc.SetStrategy(sub)
	result = calc.Calculate(x, y)
	t.Logf("sub result: %d\n", result)

	calc.SetStrategy(mul)
	result = calc.Calculate(x, y)
	t.Logf("mul result: %d\n", result)
}