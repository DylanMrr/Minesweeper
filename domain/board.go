package domain

type Board struct {
	Cells           [][]*Cell
	NotOpenMines    int
	MarkedMines     int
	NotCheckedCells int
}

func NewBoard() *Board {
	return &Board{}
}
