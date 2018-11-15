package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	rows    = 9
	columns = 9
)

var (
	ErrBounds = errors.New("out of bounds")
	ErrDigit  = errors.New("invalid digit")
)

type SudokuError []error

func (se SudokuError) Error() string {
	var errorStrings []string

	for _, e := range se {
		errorStrings = append(errorStrings, e.Error())
	}

	return strings.Join(errorStrings, ", ")
}

// Grid is a Sudoku grid
type Grid [rows][columns]int8

func (g *Grid) Set(row, column int, digit int8) error {
	var se SudokuError
	if !inBounds(row, column) {
		se = append(se, ErrBounds)
	}

	if !validDigit(digit) {
		se = append(se, ErrDigit)
	}

	if len(se) > 0 {
		return se
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

func validDigit(digit int8) bool {
	// if (digit >= 1) && (digit <= 9) {
	// 	return true
	// }
	// return false
	return (digit >= 1) && (digit <= 9)
}

func main() {
	var g Grid
	err := g.Set(9, 9, 11)
	if err != nil {
		if errs, ok := err.(SudokuError); ok {
			fmt.Printf("%d error(s) occured:\n", len(errs))
			for _, e := range errs {
				fmt.Printf("- %v\n", e)
			}
		}
		os.Exit(1)
	}
}
