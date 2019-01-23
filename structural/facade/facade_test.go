package facade

import "testing"

func TestFacade(t *testing.T) {
	computer := NewComputer()

	computer.start()
}
