package tictactui

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Board represents a 8x8 board game.
type Board struct {
	Turn   byte
	Winner byte
	Values [10]int
	Pieces [8][8]byte
	X      int
	Y      int
}

// NewBoard creates a new board.
func NewBoard(rd io.Reader) (b Board, m []string, f bool, e error) {

	scanner := bufio.NewScanner(rd)
	step    := 1
	x       := 0
	y       := 0

	/* Read header. */
	for step == 1 && scanner.Scan() {

		l := strings.Trim(scanner.Text(), " ")
		if l == "." {
			step = 3
			break
		}
		if l == "%" {
			step = 2
			break
		}

		key, val, found := strings.Cut(l, ":")
		if !found {
			continue
		}

		key = strings.Trim(key, " ")
		val = strings.Trim(val, " ")
		if len(key) == 0 {
			continue
		}
		if len(val) == 0 {
			continue
		}

		switch {
		case key == "turn":
			b.Turn = val[0]
		case key == "move":
			m = append(m, val)
		case key == "winner":
			b.Winner = val[0]
		}
	}

	/* Read table. */
	for step == 2 && scanner.Scan() {

		l := scanner.Text()
		if l[0] == '.' {
			step = 3
			break
		}

		for c := 0; c < len(l); c++ {
			if l[c] == '|' && x < 8 && y < 8 {
				if c == len(l)-1 {
					y++
					if x > b.X {
						b.X = x
					}
					x = 0
				} else {
					b.Pieces[x][y] = l[c+1]
					x++
				}
			}
		}
		b.Y = y
	}

	/* Put empty pieces. */
	for x = 0; x < b.X; x++ {
		for y = 0; y < b.Y; y++ {
			if b.Pieces[x][y] == 0 {
				b.Pieces[x][y] = ' '
			}
		}
	}

	/* Finished. */
	f = step == 3

	return
}

// Print board
func (b *Board) Print(wd io.Writer) {
	if b.Turn != 0 {
		fmt.Fprintf(wd, "turn: %c\n", b.Turn)
	}
	if b.Winner != 0 {
		fmt.Fprintf(wd, "winner: %c\n", b.Winner)
	}
	fmt.Fprintf(wd, "values: %v,%v\n", b.Values[0], b.Values[1])
	fmt.Fprintf(wd, "%%\n")
	for y := -1; y < b.Y; y++ {
		if y == -1 {
			fmt.Fprintf(wd, " ")
		} else {
			fmt.Fprintf(wd, "%c", 'a'+y)
		}
		for x := 0; x < b.X; x++ {
			if y == -1 {
				fmt.Fprintf(wd, " %c", '1'+x)
			} else {
				fmt.Fprintf(wd, "|%c", b.Pieces[x][y])
			}
		}
		if y == -1 {
			fmt.Fprintf(wd, " \n")
		} else {
			fmt.Fprintf(wd, "|\n")
		}
	}
	fmt.Fprintf(wd, ".\n")
}

// GetBoardPoint Parse "a1" like string to position.
func (b *Board) GetBoardPoint(s string) (x, y int, e error) {
	x = -1
	y = -1
	for _, c := range []byte(s) {
		if c >= 'a' && c <= 'z' {
			y = int(c - 'a')
		} else if c >= '1' && c <= '9' {
			x = int(c - '1')
		}
	}
	if x == -1 || y == -1 {
		e = fmt.Errorf("Invalid move format")
		return
	}
	if x >= b.X || y >= b.Y {
		e = fmt.Errorf("Move out of board")
		return
	}
	return
}
