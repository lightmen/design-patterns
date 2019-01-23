package decorator

import (
	"testing"
	"fmt"
)

func TestDecorator(t *testing.T) {
	var c Component = &ConcreteComponent{}
	c = NewAddDecorate(c, 10)
	c = NewMulDecorate(c, 8)

	fmt.Println("calc: ", c.Calc())
}
