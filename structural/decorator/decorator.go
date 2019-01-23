package decorator


type Component interface {
	Calc() int
}

type ConcreteComponent struct {}

func (c *ConcreteComponent)Calc() int {
	return 0
}

type MulDecorate struct {
	Component
	num int
}

func NewMulDecorate(c Component, num int) Component {
	return &MulDecorate{
		Component: c,
		num: num,
	}
}

func (d *MulDecorate)Calc() int {
	return d.Component.Calc() * d.num
}

type AddDecorate struct {
	Component
	num int
}

func NewAddDecorate(c Component, num int) Component {
	return &AddDecorate{
		Component: c,
		num: num,
	}
}

func (d *AddDecorate)Calc() int {
	return d.Component.Calc() + d.num
}