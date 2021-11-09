package game

import (
	"fmt"

	"github.com/DylanMrr/Minesweeper/core"
	"github.com/DylanMrr/Minesweeper/domain"
	"github.com/DylanMrr/Minesweeper/ui/input"
	"github.com/DylanMrr/Minesweeper/ui/output"
)

func Process() {
	board := domain.NewBoard()
	InitBoard(core.BoardSize, core.MinesCount, board)

	for board.NotCheckedCells > 0 {
		moveType, i, j := input.InputCell()

		if moveType == core.Mark {
			if !board.Cells[i][j].IsChecked {
				board.Cells[i][j].IsMarked = true
				board.Cells[i][j].IsChecked = true
				board.NotOpenMines--
				if board.Cells[i][j].State == core.Mine {
					board.MarkedMines++
				}
			} else if board.Cells[i][j].IsMarked {
				board.Cells[i][j].IsMarked = false
				board.Cells[i][j].IsChecked = false
				board.NotOpenMines++
				if board.Cells[i][j].State == core.Mine {
					board.MarkedMines--
				}
			} else {
				fmt.Println("Точка уже проверена")
			}
		} else if moveType == core.Move {
			if board.Cells[i][j].State == core.Mine {
				fmt.Println("Вы проиграли!")
				board.Cells[i][j].State = core.Lose
				output.PrintBoard(board)
				break
			} else {
				openCells(i, j, board)
			}
		} else {

		}
		if board.MarkedMines == 10 {
			fmt.Println("Вы победили!")
			break
		}
		output.PrintBoard(board)
		fmt.Println("Осталось ", board.NotOpenMines, "мин")
	}

	if board.MarkedMines == 10 {
		fmt.Println("Вы победили!")
	} else {
		fmt.Println("Вы проиграли!")
	}
}

func openCells(i, j int, board *domain.Board) {
	if i < 0 || i >= len(board.Cells) || j < 0 || j >= len(board.Cells) || board.Cells[i][j].IsChecked {
		return
	}
	board.Cells[i][j].IsChecked = true
	board.NotCheckedCells--
	if (*board).Cells[i][j].State == core.NearMine {
		return
	}
	openCells(i+1, j, board)
	openCells(i-1, j, board)
	openCells(i, j+1, board)
	openCells(i, j-1, board)
}
