package core

type CellState int

const (
	Empty CellState = iota
	NearMine
	Mine
	Lose
)
