package tictactui

import (
	"fmt"
)

// TicTacToeApply applies a move to board.
func (b *Board) TicTacToeApply(m string) (e error) {
	x, y, e := b.GetBoardPoint(m)
	if e != nil {
		return
	}
	if b.Pieces[x][y] != ' ' {
		e = fmt.Errorf("Illegal move")
		return
	}
	switch b.Turn {
	case 'o':
		b.Pieces[x][y] = 'o'
		b.Turn = 'x'
	case 'x':
		b.Pieces[x][y] = 'x'
		b.Turn = 'o'
	default:
		panic("Invalid player.")
	}
	return
}

// TicTacToeValue values a game.
func (b *Board) TicTacToeValue() {
	var pos = [8][3][2]int {
		
		{ {0,0}, {0,1}, {0,2} },
        { {1,0}, {1,1}, {1,2} },
        { {2,0}, {2,1}, {2,2} },

        { {0,0}, {1,0}, {2,0} },
        { {0,1}, {1,1}, {2,1} },
        { {0,2}, {1,2}, {2,2} },

        { {0,0}, {1,1}, {2,2} },
        { {2,0}, {1,1}, {0,2} },
	}

	for p:=0; p<8; p++ {
		p1 := b.Pieces[pos[p][0][0]][pos[p][0][1]]
        p2 := b.Pieces[pos[p][1][0]][pos[p][1][1]]
        p3 := b.Pieces[pos[p][2][0]][pos[p][2][1]]
		if (p1 != ' ' && (p1 == p2) && (p2 == p3)) {
            b.Winner     = p1;
			switch p1 {
			case 'x':
				b.Values[0] =  1;
				b.Values[1] = -1;
			case 'o':
				b.Values[0] = -1;
				b.Values[1] =  1;
			}
            return;
        }
	}

	c := 0
	for x:=0; x<b.X; x++ {
		for y:=0; y<b.Y; y++ {
            if (b.Pieces[x][y]==' ') {
                c++;
			}
		}
	}
	if (c==0) {
        b.Winner = 'n';
    }
	b.Values[0]  = 0;
	b.Values[1]  = 0;
}

// TicTacToeMoves list moves.
func (b *Board) TicTacToeMoves() (r []string) {
	if (b.Pieces[0][0] == ' ') { r = append(r, "a1") }
	if (b.Pieces[1][0] == ' ') { r = append(r, "a2") }
	if (b.Pieces[2][0] == ' ') { r = append(r, "a3") }
	if (b.Pieces[0][1] == ' ') { r = append(r, "b1") }
	if (b.Pieces[1][1] == ' ') { r = append(r, "b2") }
	if (b.Pieces[2][1] == ' ') { r = append(r, "b3") }
	if (b.Pieces[0][2] == ' ') { r = append(r, "c1") }
	if (b.Pieces[1][2] == ' ') { r = append(r, "c2") }
	if (b.Pieces[2][2] == ' ') { r = append(r, "c3") }
	return
}


// MinimaxTicTacToe ...
var MinimaxTicTacToe = Minimax[Board] {
	GetMoves  : func (b *Board) (r []string) {
		return b.TicTacToeMoves()
	},
	GetValue  : func (b *Board, player byte) (value int, turn byte, winner byte) {
		b.TicTacToeValue();
		switch player {
		case 'x': value = b.Values[0]
		case 'o': value = b.Values[1]
		}
		winner = b.Winner
		turn   = b.Turn
		return
	},
	ApplyMove : func (b *Board, mv string) (r Board, e error) {
		r = *b;
		r.TicTacToeApply(mv)
		return
	},
}
