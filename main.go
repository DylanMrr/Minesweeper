package main

import (
	"fmt"

	"github.com/DylanMrr/Minesweeper/domain"
	"github.com/DylanMrr/Minesweeper/game"
)

func main() {
	fmt.Println("mines")
	game.InitBoard(10, 10, domain.NewBoard())
}
