package core

type CellState int

const (
	Empty CellState = iota
	NearMine
	Mine
	Checked
	Lose
)
