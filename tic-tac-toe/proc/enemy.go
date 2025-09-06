package proc

import (
	"math/rand"

	"github.com/logica0419/tinygo-keeb/tic-tac-toe/types"
)

func (p *Processor) enemyInput() {
	if p.diff == types.Easy {
		p.enemyRand()
		return
	}

	if p.enemyToWin() {
		return
	}

	if p.diff == types.Hard {
		if p.enemyNotToLose() {
			return
		}
		if p.enemyRiichi() {
			return
		}
	}

	if p.enemyMiddle() {
		return
	}

	p.enemyRand()
}

func (p *Processor) enemyNotToLose() bool {
	for _, v := range judgePairs {
		if p.st[v[0]] == types.Circle && p.st[v[1]] == types.Circle && p.st[v[2]] != types.Cross {
			p.st[v[2]] = types.Cross
			return true
		}

		if p.st[v[1]] == types.Circle && p.st[v[2]] == types.Circle && p.st[v[0]] != types.Cross {
			p.st[v[0]] = types.Cross
			return true
		}

		if p.st[v[0]] == types.Circle && p.st[v[2]] == types.Circle && p.st[v[1]] != types.Cross {
			p.st[v[1]] = types.Cross
			return true
		}
	}

	return false
}

func (p *Processor) enemyToWin() bool {
	for _, v := range judgePairs {
		if p.st[v[0]] == types.Cross && p.st[v[1]] == types.Cross && p.st[v[2]] != types.Circle {
			p.st[v[2]] = types.Cross
			return true
		}

		if p.st[v[1]] == types.Cross && p.st[v[2]] == types.Cross && p.st[v[0]] != types.Circle {
			p.st[v[0]] = types.Cross
			return true
		}

		if p.st[v[0]] == types.Cross && p.st[v[2]] == types.Cross && p.st[v[1]] != types.Circle {
			p.st[v[1]] = types.Cross
			return true
		}
	}

	return false
}

func (p *Processor) enemyRiichi() bool {
	for _, v := range judgePairs {
		if p.st[v[0]] == types.Cross {
			if p.st[v[1]] != types.Cross && p.st[v[1]] != types.Circle && p.st[v[2]] != types.Cross && p.st[v[2]] != types.Circle {
				p.st[v[rand.Intn(2)+1]] = types.Cross
				return true
			}
		}

		if p.st[v[1]] == types.Cross {
			if p.st[v[0]] != types.Cross && p.st[v[0]] != types.Circle && p.st[v[2]] != types.Cross && p.st[v[2]] != types.Circle {
				if rand.Intn(2) == 0 {
					p.st[v[0]] = types.Cross
				} else {
					p.st[v[2]] = types.Cross
				}
				return true
			}
		}

		if p.st[v[2]] == types.Cross {
			if p.st[v[0]] != types.Cross && p.st[v[0]] != types.Circle && p.st[v[1]] != types.Cross && p.st[v[1]] != types.Circle {
				p.st[v[rand.Intn(2)]] = types.Cross
				return true
			}
		}
	}

	return false
}

func (p *Processor) enemyMiddle() bool {
	if p.st[4] != types.Circle && p.st[4] != types.Cross {
		p.st[4] = types.Cross
		return true
	}

	return false
}

func (p *Processor) enemyRand() {
	validCell := []int{}
	for i := range p.st {
		if p.st[i] == types.Empty {
			validCell = append(validCell, i)
		}
	}

	modCell := validCell[rand.Intn(len(validCell))]
	p.st[modCell] = types.Cross
}
