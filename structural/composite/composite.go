package composite

import "fmt"

type Component interface {
	Add(component Component)
	Remove(component Component)
	Print(pre string)
}


type DefaultComponent struct {
	Name string
}

func (d *DefaultComponent)Add(component Component){

}

func (d *DefaultComponent)Remove(component Component) {

}

func (d *DefaultComponent)Print(pre string){

}

type Leaf struct {
	*DefaultComponent
}

func NewLeaf(name string) *Leaf {
	return &Leaf{
		DefaultComponent: &DefaultComponent{
			Name: name,
		},
	}
}

func (l *Leaf)Print(pre string) {
	fmt.Printf("%s %s\n", pre, l.Name)
}


type Composite struct {
	*DefaultComponent
	components []Component
}

func NewComposite(name string) *Composite{
	return &Composite{
		DefaultComponent: &DefaultComponent{
			Name: name,
		},
		components: make([]Component, 0),
	}
}

func (c *Composite)Add(component Component) {
	c.components = append(c.components, component)
}

func (c *Composite)Remove(component Component) {
	for idx, item := range c.components {
		if item == component {
			c.components = append(c.components[0:idx], c.components[idx+1:]...)
			break;
		}
	}
}

func (c *Composite)Print(pre string) {
	fmt.Printf("%s%s: \n", pre, c.Name)
	for _, item := range c.components {
		item.Print(pre + "-")
	}
}