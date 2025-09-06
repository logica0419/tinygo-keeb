package proc

import (
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/logica0419/tinygo-keeb/tic-tac-toe/display"
	"github.com/logica0419/tinygo-keeb/tic-tac-toe/machine"
	"github.com/logica0419/tinygo-keeb/tic-tac-toe/types"
)

type Processor struct {
	diff int
	turn types.Turn
	st   [9]types.State

	disp *display.Display
	mac  *machine.WS2812B
}

func New() *Processor {
	mac := machine.NewWS2812B()
	disp := display.New(mac)

	return &Processor{
		st: [9]types.State{
			types.Empty, types.Empty, types.Empty,
			types.Empty, types.Empty, types.Empty,
			types.Empty, types.Empty, types.Empty,
		},

		disp: disp,
		mac:  mac,
	}
}

func (p *Processor) Game() error {
	p.diff = types.Hard
	p.turn = rand.IntN(2) == 0

	p.mac.DisplayString("Running...")

	for i := 0; i < 9; i++ {
		err := p.disp.State(p.st, p.turn)
		if err != nil {
			return fmt.Errorf("failed to update display: %w", err)
		}

		if p.turn {
			place := p.mac.Scan()
			err := p.input(place, types.Circle)
			if err != nil {
				i--
				continue
			}
		} else {
			place := p.mac.Scan()
			err := p.input(place, types.Cross)
			if err != nil {
				i--
				continue
			}
		}

		res := p.judge()
		if res != types.Empty {
			break
		}

		p.turn = !p.turn

		time.Sleep(10 * time.Millisecond)
	}

	err := p.disp.State(p.st, p.turn)
	if err != nil {
		return fmt.Errorf("failed to update display: %w", err)
	}

	switch p.judge() {
	case types.Circle:
		p.mac.DisplayString("Red Win!")
	case types.Cross:
		p.mac.DisplayString("Blue Win!")
	case types.Empty:
		p.mac.DisplayString("Draw")
	}

	time.Sleep(2 * time.Second)

	p.mac.DisplayString("Press Any\nKey to\nRestart")
	_ = p.mac.Scan()

	return nil
}

func (p *Processor) Reset() {
	p.st = [9]types.State{
		types.Empty, types.Empty, types.Empty,
		types.Empty, types.Empty, types.Empty,
		types.Empty, types.Empty, types.Empty,
	}
}
