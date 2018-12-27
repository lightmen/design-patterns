package template_method

import "fmt"

type ITemplate interface {
	TemplateMethod()
}

type template struct {
	implement
}

func (this *template) TemplateMethod() {
	this.PrivateOperation1()
	this.PrivateOperation2()
}

func newTemplate(impl implement) *template {
	return &template{
		implement: impl,
	}
}
type implement interface {
	PrivateOperation1()
	PrivateOperation2()
}

type ContreteImplement1 struct {
	*template
}

func NewContreteImplement1() ITemplate{
	c := &ContreteImplement1{}
	template := newTemplate(c)
	c.template = template
	return c
}

func (this *ContreteImplement1) PrivateOperation1() {
	fmt.Println("Contrete1 - PrivateOperation1")
}

func (this *ContreteImplement1) PrivateOperation2() {
	fmt.Println("Contrete1 - PrivateOperation2")
}

type ContreteImplement2 struct {
	*template
}

func NewContreteImplement2() ITemplate{
	c := &ContreteImplement2{}
	template := newTemplate(c)
	c.template = template
	return c
}

func (this *ContreteImplement2) PrivateOperation1() {
	fmt.Println("Contrete2 - PrivateOperation1")
}

func (this *ContreteImplement2) PrivateOperation2() {
	fmt.Println("Contrete2 - PrivateOperation2")
}
