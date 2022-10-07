package tictactui

import (
	"math"
)

// Minimax functions used by the algorithm.
type Minimax[State any] struct {
	GetMoves  func (s *State)              []string
	GetValue  func (s *State, player byte) (value int, turn byte, winner byte)
	ApplyMove func (s *State, move string) (State, error)
}

// MinimaxRun the minimax algorithm to get the best move.
func MinimaxRun[State any] (m *Minimax[State], s1 *State, ia byte, maxDepth int) (value int, move string) {

	value, turn, winner := m.GetValue(s1, ia)
	maximizer := ia == turn
	moves     := m.GetMoves(s1)
	
	if (winner!=0 || maxDepth == 0 || len(moves) == 0) {
		
	} else {
		if maximizer {
			value = math.MinInt
		} else {
			value = math.MaxInt
		}
		for _,mv := range moves {
			s2, e := m.ApplyMove (s1, mv)
			if e != nil { continue }
			nvalue, _ := MinimaxRun[State](m, &s2, ia, maxDepth-1)
			if ((maximizer == true  && nvalue > value) ||
				(maximizer == false && nvalue < value)) {
				value = nvalue
				move = mv
			}
		}
	}
	
	return
}
