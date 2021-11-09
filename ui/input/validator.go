package input

import (
	"errors"
	"strconv"
	"strings"

	"github.com/DylanMrr/Minesweeper/core"
)

func ValidateCell(cell *string) (core.MoveType, int, int, error) {
	elements := strings.Split(*cell, " ")
	if len(elements) != 3 {
		return -1, -1, -1, errors.New("Некорректное количество введенных элементов")
	}

	t, err := strconv.Atoi(elements[0])
	if err != nil || t != 0 && t != 1 {
		return -1, -1, -1, errors.New("Некорректный первый индекс")
	}
	moveType := core.MoveType(t)

	i, err := strconv.Atoi(elements[1])
	if err != nil || i < 1 || i > core.BoardSize {
		return -1, -1, -1, errors.New("Некорректный второй индекс")
	}

	last := strings.Trim(elements[2], "\r\t\n")
	j, err := strconv.Atoi(last)
	if err != nil || j < 1 || j > core.BoardSize {
		return -1, -1, -1, errors.New("Некорректный третий индекс")
	}

	return moveType, i - 1, j - 1, nil
}
