package game

import (
	"math/rand"
	"time"

	"github.com/DylanMrr/Minesweeper/core"
	"github.com/DylanMrr/Minesweeper/domain"
	"github.com/DylanMrr/Minesweeper/ui/output"
)

func InitBoard(size, minesCount int, board *domain.Board) {
	(*board).Cells = make([][]*domain.Cell, size)
	for i := 0; i < size; i++ {
		(*board).Cells[i] = make([]*domain.Cell, size)
		for j := 0; j < size; j++ {
			(*board).Cells[i][j] = domain.NewCell(core.Empty, 0)
		}
	}
	(*board).NotOpenMines = minesCount
	(*board).NotCheckedCells = size * size
	rand.Seed(time.Now().UnixNano())
	for k := 0; k < minesCount; k++ {
		f := true
		for f {
			i := rand.Intn(size)
			j := rand.Intn(size)
			if (*board).Cells[i][j].State == core.Empty {
				(*board).Cells[i][j].State = core.Mine
				f = false
			}
		}
	}

	for i := range board.Cells {
		for j := range board.Cells[i] {
			if (*board).Cells[i][j].State == core.Mine {
				markNearMine(i+1, j, board)
				markNearMine(i-1, j, board)
				markNearMine(i, j+1, board)
				markNearMine(i, j-1, board)

				markNearMine(i-1, j-1, board)
				markNearMine(i+1, j-1, board)
				markNearMine(i-1, j+1, board)
				markNearMine(i+1, j+1, board)
			}
		}
	}

	output.PrintBoard(board)
}

func markNearMine(i, j int, board *domain.Board) {
	if i < 0 || i >= len(board.Cells) || j < 0 || j >= len(board.Cells) || (*board).Cells[i][j].State == core.Mine {
		return
	}
	(*board).Cells[i][j].State = core.NearMine
	(*board).Cells[i][j].MinesCount++
}
