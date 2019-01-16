package flyweight

import "fmt"

type FlyweightFactory struct {
	mpFlyweight map[string] Flyweight
}

func NewFlyweightFactory() *FlyweightFactory{
	return &FlyweightFactory{
		mpFlyweight: make(map[string]Flyweight),
	}
}

func (f *FlyweightFactory)getFlyweight(name string) Flyweight {
	fw, ok := f.mpFlyweight[name]
	if ok {
		return fw
	}

	fw = NewConcreteFlyweight(name)
	f.mpFlyweight[name] = fw
	return fw
}

type Flyweight interface {
	operation()
}

type ConcreteFlyweight struct {
	name string
}

func NewConcreteFlyweight(name string) *ConcreteFlyweight {
	return &ConcreteFlyweight{
		name: name,
	}
}

func (f *ConcreteFlyweight)operation() {
	fmt.Printf("ConcreteFlyweight: %s\n", f.name)
}

