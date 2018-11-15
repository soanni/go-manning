package main

import (
	"errors"
	"fmt"
	"os"
)

const (
	rows    = 9
	columns = 9
)

var (
	ErrBounds = errors.New("out of bounds")
	ErrDigit  = errors.New("invalid digit")
)

// Grid is a Sudoku grid
type Grid [rows][columns]int8

func (g *Grid) Set(row, column int, digit int8) error {
	if !inBounds(row, column) {
		return ErrBounds
	}

	g[row][column] = digit
	return nil
}

func inBounds(row, column int) bool {
	if row < 0 || row >= rows {
		return false
	}
	if column < 0 || column >= columns {
		return false
	}
	return true
}

func main() {
	var g Grid
	err := g.Set(10, 0, 5)
	fmt.Printf("%T\n",ErrBounds)
	fmt.Printf("%T\n",ErrDigit)
	if err != nil {
		switch err {
		case ErrBounds, ErrDigit:
			fmt.Println("Les erreurs de paramètres hors limites.")
		default:
			fmt.Println(err)
		}
		os.Exit(1)
	}
}
