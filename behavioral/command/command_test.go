package command

import "testing"

func TestCommand(t *testing.T) {
	switcher := NewSwitcher()

	light := NewLight()

	onCmd := NewOnCommand(light)
	offCmd := NewOffCommand(light)

	switcher.StoreAndExecute(onCmd)
	switcher.StoreAndExecute(offCmd)
}
