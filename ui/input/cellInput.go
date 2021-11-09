package input

import (
	"bufio"
	"fmt"
	"os"

	"github.com/DylanMrr/Minesweeper/core"
)

func InputCell() (core.MoveType, int, int) {
	isValid := false

	cin := bufio.NewReader(os.Stdin)

	for !isValid {
		fmt.Println("Введите точку для стрельбы в формате 0 10 5, где")
		fmt.Println("Первая позиция: 0 - открыть ячейку, 1 - пометить как мину")
		fmt.Println("Вторая позиция: позиция по вертикали")
		fmt.Println("Третья позиция: позиция по горизонтали")

		cell, _ := cin.ReadString('\n')

		moveType, i, j, err := ValidateCell(&cell)
		if err != nil {
			fmt.Println(err)
			continue
		}

		return moveType, i, j
	}
	return -1, -1, -1
}
