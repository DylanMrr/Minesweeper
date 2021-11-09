package game

import (
	"fmt"
	"strconv"

	"github.com/DylanMrr/Minesweeper/core"
)

func InitConfig() {
	for true {
		fmt.Println("Выберите сложность")
		fmt.Println("1 - 7х7, 7 мин")
		fmt.Println("2 - 10х10, 15 мин")
		fmt.Println("3 - 15х15, 30 мин")
		fmt.Println("Введите цифру - сложность")
		var lvl string
		fmt.Scan(&lvl)

		i, err := strconv.Atoi(lvl)
		if err != nil {
			fmt.Println("Некорректный уровень")
			continue
		}
		if i < 1 || 1 > 3 {
			fmt.Println("Некорректный уровень")
			continue
		}
		switch i {
		case 1:
			core.BoardSize = 7
			core.MinesCount = 7
		case 2:
			core.BoardSize = 10
			core.MinesCount = 15
		case 3:
			core.BoardSize = 15
			core.MinesCount = 30
		}
		break
	}
}
