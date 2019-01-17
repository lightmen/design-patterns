package builder

type Director struct {
	pbuilder Builder
}

func NewDirector() *Director {
	return &Director{}
}

func (d *Director)SetBuilder(builder Builder) {
	d.pbuilder = builder
}

func (d *Director)Construct() {
	d.pbuilder.BuildPartA()
	d.pbuilder.BuildPartB()
	d.pbuilder.BuildPartC()
}

type Builder interface {
	BuildPartA()
	BuildPartB()
	BuildPartC()
}


type ConcreteBuilder struct {
	step int
}

func NewConcreteBuilder() *ConcreteBuilder{
	return &ConcreteBuilder{}
}


func (cb *ConcreteBuilder)BuildPartA() {
	cb.step = 1
}

func (cb *ConcreteBuilder)BuildPartB() {
	cb.step += 1
}


func (cb *ConcreteBuilder)BuildPartC() {
	cb.step += 1
}