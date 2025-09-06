package types

type State int

const (
	Empty State = iota
	Circle
	Cross
)

const (
	Easy = iota
	Normal
	Hard
)

type Turn bool

const (
	Player Turn = true
	Enemy  Turn = false
)
