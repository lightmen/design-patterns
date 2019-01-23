package facade

import "fmt"

type CPU struct {}

func (c *CPU)execute() {
	fmt.Println("cpu execute")
}

type Memory struct {}

func (m *Memory)load() {
	fmt.Println("memory load")
}

type Hard struct{}

func (h *Hard) read() {
	fmt.Println("hard read")
}

type Computer struct {
	cpu *CPU
	memory *Memory
	hard *Hard
}

func NewComputer() *Computer {
	return &Computer{
		cpu: &CPU{},
		memory: &Memory{},
		hard: &Hard{},
	}
}

func (c *Computer)start() {
	c.cpu.execute()
	c.memory.load()
	c.hard.read()
}
