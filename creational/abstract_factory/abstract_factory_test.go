package abstract_factory

import (
	"testing"
	"fmt"
)

func TestAbstractFactory(t *testing.T) {
	var factory AbstractFactory
	var button Button
	var window Window

	factory = &WinFactory{}
	button = factory.CreateButton()
	window = factory.CreateWindow()

	fmt.Printf("button: %s, window: %s\n", button.Press(), window.View())

	factory = &MacFactory{}
	button = factory.CreateButton()
	window = factory.CreateWindow()

	fmt.Printf("button: %s, window: %s\n", button.Press(), window.View())
}
