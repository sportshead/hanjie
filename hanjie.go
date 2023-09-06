package main

import (
	"fmt"
	"strings"
	"time"
)

type Status int

const (
	StatusUnset Status = iota
	StatusFilled
	StatusEmpty
)

type Hint []int

type Hanjie struct {
	Width  int
	Height int
	Grid   [][]Status
	Hints  struct {
		Row    []Hint
		Column []Hint

		MaxRowSize *int
		MaxColSize *int
	}
	StepDelay time.Duration
}

func (h *Hanjie) Print(message ...any) {
	fmt.Println(message...)

	fmt.Println("Grid:")
	offset := 7 + *h.Hints.MaxRowSize*3
	padding := strings.Repeat(" ", offset)

	for i := *h.Hints.MaxColSize - 1; i >= 0; i-- {
		fmt.Print(padding, "|")
		for _, col := range h.Hints.Column {
			if i >= len(col) {
				fmt.Print("  |")
			} else {
				fmt.Printf("%2d|", col[len(col)-1-i])
			}
		}
		fmt.Println()
	}

	for i, row := range h.Grid {
		fmt.Printf("Row %02d: %s| ", i+1, hintToString(h.Hints.Row[i], *h.Hints.MaxRowSize))
		fmt.Printf("%s|\n", rowToString(row))
	}
	time.Sleep(h.StepDelay)
}

func (h *Hanjie) InitializeGrid() {
	h.Grid = make([][]Status, h.Height)
	for i := 0; i < h.Height; i++ {
		h.Grid[i] = make([]Status, h.Width)
		for j := 0; j < h.Width; j++ {
			h.Grid[i][j] = StatusUnset
		}
	}
}

// runs all solving algorithms
func (h *Hanjie) SolveAll() {
	h.overlap()
}

// https://en.wikipedia.org/wiki/Nonogram#Simple_boxes
func (h *Hanjie) overlap() {
	step := "overlap() Rows"
	for i, hints := range h.Hints.Row {
		rowLeft := 0
		for j, hint := range hints {
			rightBound := rowLeft + hint
			leftBound := h.Width - hint
			for k := len(hints) - 1; k > j; k-- {
				leftBound -= hints[k] + 1
			}

			fmt.Println(leftBound, rightBound)
			for k := leftBound; k <= rightBound; k++ {
				h.Grid[i][k-1] = StatusFilled
			}

			rowLeft += hint + 1
		}
		h.Print(step, i+1)
	}
	step = "overlap() cols"
	for i, hints := range h.Hints.Column {
		colTop := 0
		for j, hint := range hints {
			bottomBound := colTop + hint
			topBound := h.Height - hint
			for k := len(hints) - 1; k > j; k-- {
				topBound -= hints[k] + 1
			}

			for k := topBound; k <= bottomBound; k++ {
				h.Grid[k-1][i] = StatusFilled
			}

			colTop += hint + 1
		}
		h.Print(step, i+1)
	}
}

// https://en.wikipedia.org/wiki/Nonogram#Simple_spaces
func (h *Hanjie) simpleSpaces() {
	step := "simpleSpaces() Rows"
	for i, hints := range h.Hints.Row {
		_ = hints

		h.Print(step, i+1)
	}
	step = "simpleSpaces() cols"
	for i, hints := range h.Hints.Column {
		_ = hints

		h.Print(step, i+1)
	}
}

func rowToString(row []Status) string {
	rowStr := make([]string, len(row))
	for i, status := range row {
		rowStr[i] = statusToString(status)
	}
	return strings.Join(rowStr, "| ")
}

func statusToString(status Status) string {
	switch status {
	case StatusFilled:
		return "■"
	case StatusEmpty:
		return "□"
	default:
		return " "
	}
}

func hintToString(hint Hint, length int) string {
	hintStr := make([]string, length)
	pad := length - len(hint)
	for i := 0; i < pad; i++ {
		hintStr[i] = "  "
	}
	for i := pad; i < length; i++ {
		hintStr[i] = fmt.Sprintf("%2d", hint[i-pad])
	}
	return strings.Join(hintStr, " ")
}
