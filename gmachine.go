// Package gmachine implements a simple virtual CPU, known as the G-machine.
package gmachine

// DefaultMemSize is the number of 64-bit words of memory which will be
// allocated to a new G-machine by default.
const DefaultMemSize = 1024

type Machine struct{
	Memory []uint64
	P uint64
}

func New() Machine {
	return Machine{
		Memory: make([]uint64, DefaultMemSize),
	}
}

func (g *Machine) Run() {
	for {
		op := g.Memory[g.P]
		g.P++
		switch op {
		case 0:
			return
		case 1:
		}
	}
}