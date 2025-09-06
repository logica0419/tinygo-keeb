package display

import (
	"github.com/logica0419/tinygo-keeb/tic-tac-toe/machine"
	"github.com/logica0419/tinygo-keeb/tic-tac-toe/types"
)

const (
	black uint32 = 0x00000000
	white uint32 = 0xFFFFFFFF
	red   uint32 = 0x00FF00FF
	blue  uint32 = 0x0000FFFF
)

var displayMap = [9]int{0, 3, 6, 1, 4, 7, 2, 5, 8}

type Display struct {
	ws *machine.WS2812B
}

func New(ws *machine.WS2812B) *Display {
	return &Display{
		ws: ws,
	}
}

func (d *Display) State(st [9]types.State, turn types.Turn) error {
	colors := []uint32{
		black, black, black, black,
		black, black, black, black,
		black, black, black, black,
	}
	for i, v := range st {
		switch v {
		case types.Empty:
			colors[displayMap[i]] = black
		case types.Circle:
			colors[displayMap[i]] = red
		case types.Cross:
			colors[displayMap[i]] = blue
		}
	}

	switch turn {
	case types.Player:
		colors[11] = white
	case types.Enemy:
		colors[9] = white
	}

	return d.ws.WriteRaw(colors)
}
