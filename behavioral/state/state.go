package state

import "fmt"

type State interface {
	Handle(player *MP3Player)
}

type Standby struct {}

func (s *Standby)Handle(player *MP3Player) {
	fmt.Println("standby")
	player.SetState(new(Playing))
}

type Playing struct{}

func (p *Playing)Handle(player *MP3Player) {
	fmt.Println("playing")
	player.SetState(new(Standby))
}

type MP3Player struct {
	state State
}

func NewMP3Player(state State) *MP3Player {
	return &MP3Player{
		state: state,
	}
}

func (p *MP3Player)SetState(state State) {
	p.state = state
}

func (p *MP3Player)Play() {
	p.state.Handle(p)
}