package factory_method

import (
	"testing"
	"fmt"
)

func TestFactoryMethod(t *testing.T) {
	wf := &WinButtonFactory{}

	btn := wf.CreateButton()
	fmt.Println(btn.Show())

	mf := &MacButtonFactory{}
	btn = mf.CreateButton()
	fmt.Println(btn.Show())
}
