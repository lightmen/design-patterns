package bridge

import "testing"

func TestBridge(t *testing.T) {
	api1 := &DrawAPI1{}
	shape := NewCircleShape(1, 2, 3, api1)
	shape.Draw()
}
