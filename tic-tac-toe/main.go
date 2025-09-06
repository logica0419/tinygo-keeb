package main

import (
	"github.com/logica0419/tinygo-keeb/tic-tac-toe/proc"
)

func main() {
	p := proc.New()

	for {
		err := p.Game()
		if err != nil {
			panic(err)
		}

		p.Reset()
	}
}
