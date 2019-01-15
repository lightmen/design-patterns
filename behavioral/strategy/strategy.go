package strategy

type Strategy interface {
	Execute(int, int) int
}

type Add struct{}
func (a *Add)Execute(x, y int) int {
	return x + y
}

type Subtrack struct{}
func (s *Subtrack)Execute(x, y int) int {
	return x - y
}

type Multiply struct{}
func (m *Multiply)Execute(x, y int) int {
	return x * y
}

type Calculator struct {
	strategy Strategy
}

func (c *Calculator)SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Calculator)Calculate(x, y int) int {
	return c.strategy.Execute(x, y)
}
