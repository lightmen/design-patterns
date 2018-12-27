package template_method

import "testing"

func TestContrete1(t *testing.T){
	c1 := NewContreteImplement1()

	c1.TemplateMethod()
}


func TestContrete2(t *testing.T){
	c2 := NewContreteImplement2()

	c2.TemplateMethod()
}