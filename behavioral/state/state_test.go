package state

import "testing"

func TestState(t *testing.T) {
	player := NewMP3Player(new(Standby))

	player.Play()
	player.Play()
}
