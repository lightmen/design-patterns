package flyweight

import "testing"

func TestFlyweight(t *testing.T) {
	factor := NewFlyweightFactory()

	f1 := factor.getFlyweight("f1");
	f1.operation()

	f2 := factor.getFlyweight("f2");
	f2.operation()

	f1 = factor.getFlyweight("f1");
	f1.operation()
}
