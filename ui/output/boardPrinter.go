package output

import (
	"fmt"

	"github.com/DylanMrr/Minesweeper/core"
	"github.com/DylanMrr/Minesweeper/domain"
)

func PrintBoard(board *domain.Board) {
	fmt.Print("  |")
	for i := range board.Cells {
		fmt.Print((i + 1), " ")
	}
	fmt.Println()
	fmt.Print("__")
	for range board.Cells {
		fmt.Print("__")
	}
	fmt.Println()
	for i := 0; i < len((*board).Cells); i++ {
		fmt.Print(i + 1)
		if i+1 < 10 {
			fmt.Print(" ")
		}
		fmt.Print("|")
		for j := 0; j < len((*board).Cells[i]); j++ {
			if (*board).Cells[i][j].State == core.Lose {
				fmt.Print("x ")
				continue
			}
			if (*board).Cells[i][j].IsMarked {
				fmt.Print("! ")
			} else if (*board).Cells[i][j].IsChecked {
				if (*board).Cells[i][j].State == core.NearMine {
					fmt.Print((*board).Cells[i][j].MinesCount, " ")
				} else {
					fmt.Print("  ")
				}
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
