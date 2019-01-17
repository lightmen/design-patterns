package builder

import "testing"

func TestBuild(t *testing.T) {
	builder := NewConcreteBuilder()

	director := NewDirector()
	director.SetBuilder(builder)
	director.Construct()

	if builder.step != 3 {
		t.Fail()
	}
}
