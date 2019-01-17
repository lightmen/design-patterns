package factory_method

import (
	"testing"
	"fmt"
)

func TestFactoryMethod(t *testing.T) {
	var factory ButtonFactory
	var btn Button
	factory = &WinButtonFactory{}

	btn = factory.CreateButton()
	fmt.Println(btn.Show())

	factory = &WinButtonFactory{}
	btn = factory.CreateButton()
	fmt.Println(btn.Show())
}
