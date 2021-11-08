package domain

type Board struct {
	Cells [][]*Cell
}

func NewBoard() *Board {
	return &Board{}
}
