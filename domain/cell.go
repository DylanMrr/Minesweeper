package domain

import "github.com/DylanMrr/Minesweeper/core"

type Cell struct {
	State      core.CellState
	MinesCount int
}

func NewCell(state core.CellState, minesCount int) *Cell {
	return &Cell{State: state, MinesCount: minesCount}
}
