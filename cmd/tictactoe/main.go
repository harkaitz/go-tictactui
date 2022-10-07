package main

import (
	"fmt"
	"io"
	"os"
	"log"
	t "github.com/harkaitz/go-tictactui"
	"strings"
	"github.com/pborman/getopt/v2"
)


const help string =
`Usage: tictactoe OPTS...

The classic TicTacToe game that follows the UNIX
philosophy written in Go.

-g      : Create a new game, instead of reading from the
          standard input.
-m MOVE : Specify a move, ie 'a1'.
-i      : Select a move using the Minimax algorithm.
-a      : Apply the specified moves.`
const copyrightLine string =
`Bug reports, feature requests to gemini|https://harkadev.com/oss
Copyright (c) 2022 Harkaitz Agirre, harkaitz.aguirre@gmail.com`
const board string = `
turn: x
%
  1 2 3
a| | | |
b| | | |
c| | | |
.`;


func main() {

	hFlag := getopt.BoolLong ("help", 'h')
	gFlag := getopt.Bool     ('g')
	mFlag := getopt.List     ('m', "")
	iFlag := getopt.Bool     ('i')
	aFlag := getopt.Bool     ('a')
	getopt.SetUsage(func() { fmt.Println(help + "\n\n" + copyrightLine) })
	getopt.Parse()

	if *hFlag {
		getopt.Usage()
		return
	}

	var input io.Reader
	if *gFlag {
		input = strings.NewReader(board)
	} else {
		input = os.Stdin
	}
	for {
		board, moves, found, err := t.NewBoard(input)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		if !found {
			break
		}
		if board.X != 3 || board.Y != 3 {
			log.Fatal("Invalid board size.")
			os.Exit(1)
		}
		moves = append(moves, *mFlag...)
		if *iFlag {
			_, mv := t.MinimaxRun[t.Board](&t.MinimaxTicTacToe, &board, board.Turn, 10)
			moves = append(moves, mv)
		}
		if *aFlag {
			for _,mv := range moves {
				if mv == "ai" {
					_, mv = t.MinimaxRun[t.Board](&t.MinimaxTicTacToe, &board, board.Turn, 10);
				}
				err := board.TicTacToeApply(mv)
				if err != nil {
					log.Fatal(err)
					os.Exit(1)
				}
			}
			moves = []string{}
		}
		board.TicTacToeValue()
		board.Print(os.Stdout)
	}
}



