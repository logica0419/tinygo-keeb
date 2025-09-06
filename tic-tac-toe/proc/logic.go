package proc

import (
	"fmt"

	"github.com/logica0419/tinygo-keeb/tic-tac-toe/types"
)

var judgePairs = [8][3]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {0, 3, 6}, {1, 4, 7}, {2, 5, 8}, {0, 4, 8}, {2, 4, 6}}

func (p *Processor) judge() types.State {
	for _, v := range judgePairs {
		if p.st[v[0]] == p.st[v[1]] && p.st[v[1]] == p.st[v[2]] {
			return p.st[v[0]]
		}
	}

	return types.Empty
}

func (p *Processor) input(pos int, s types.State) error {
	if pos < 0 || pos > 8 {
		return fmt.Errorf("position must be between 0 and 8")
	}

	if p.st[pos] != types.Empty {
		return fmt.Errorf("position %d is already occupied", pos)
	}

	p.st[pos] = s
	return nil
}
