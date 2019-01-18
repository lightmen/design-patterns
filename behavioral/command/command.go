package command

import "fmt"

type Light struct {
	isOn bool
}

func NewLight() *Light {
	return &Light{}
}

func (l *Light)print() {
	fmt.Println("Light status: ", l.isOn)
}

func (l *Light)TurnOn() {
	l.isOn = true
	l.print()
}

func (l *Light)TurnOff() {
	l.isOn = false
	l.print()
}


type Command interface {
	Execute()
}

type OnCommand struct{
	light *Light
}

func NewOnCommand(light *Light) *OnCommand {
	return &OnCommand{
		light: light,
	}
}

func (c *OnCommand)Execute() {
	c.light.TurnOn()
}

type OffCommand struct{
	light *Light
}

func NewOffCommand(light *Light) *OffCommand {
	return &OffCommand{
		light: light,
	}
}

func (c *OffCommand)Execute() {
	c.light.TurnOff()
}


type Switcher struct {
	cmdList []Command
}

func NewSwitcher() *Switcher {
	return &Switcher{
		cmdList: make([]Command, 0),
	}
}

func (s *Switcher)StoreAndExecute(cmd Command) {
	s.cmdList = append(s.cmdList, cmd)
	cmd.Execute()
}