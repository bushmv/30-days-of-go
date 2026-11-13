package main

import (
	"errors"
	"math"
	"strings"
)

const (
	EMPTY = '.'
	X     = 'X'
	O     = 'O'
)

const (
	stateRunning = iota
	stateDraw
	stateWon
	stateLose
)

type TicTacToe struct {
	field          [3][3]rune
	humanSymbol    rune
	computerSymbol rune
	state          int
}

type move struct {
	y int
	x int
}

func NewTicTacToe(moveFirst bool) *TicTacToe {
	game := &TicTacToe{
		state: stateRunning,
	}
	game.field = [3][3]rune{}
	for y := range 3 {
		for x := range 3 {
			game.field[y][x] = EMPTY
		}
	}
	game.humanSymbol = X
	game.computerSymbol = O
	if !moveFirst {
		game.humanSymbol = O
		game.computerSymbol = X
		// make first move
	}
	return game
}

func (t *TicTacToe) move(y, x int) error {
	if t.state != stateRunning {
		return errors.New("game already ended")
	}
	if y < 0 || y >= 3 || x < 0 || x >= 3 {
		return errors.New("move out of field")
	}
	if t.field[y][x] != EMPTY {
		return errors.New("you cannot overwrite moves")
	}
	t.field[y][x] = t.humanSymbol
	t.updateState()
	if t.state == stateRunning {
		t.computerMove()
		t.updateState()
	}
	return nil
}

func (t *TicTacToe) updateState() {
	if isWinner(t.field, t.computerSymbol) {
		t.state = stateLose
	} else if isWinner(t.field, t.humanSymbol) {
		t.state = stateWon
	} else if len(t.avaliableMoves()) == 0 {
		t.state = stateDraw
	}
}

func (t *TicTacToe) computerMove() {
	bestScore := math.Inf(-1)
	var bestMove move

	for _, move := range t.avaliableMoves() {
		t.field[move.y][move.x] = t.computerSymbol
		score := t.minmax(1, false)
		if score > bestScore {
			bestScore = score
			bestMove = move
		}
		t.field[move.y][move.x] = EMPTY
	}
	t.field[bestMove.y][bestMove.x] = t.computerSymbol
}

func (t *TicTacToe) minmax(deep int, maximize bool) float64 {
	if t.isComputerWon() {
		return 1 / float64(deep)
	} else if t.isHumanWon() {
		return -1 / float64(deep)
	}
	moves := t.avaliableMoves()
	if len(moves) == 0 {
		return 0
	}
	var best float64
	if maximize {
		best = math.Inf(-1)
		for _, move := range moves {
			t.field[move.y][move.x] = t.computerSymbol
			score := t.minmax(deep+1, !maximize)
			best = max(best, score)
			t.field[move.y][move.x] = EMPTY
		}
	} else {
		best = math.Inf(1)
		for _, move := range moves {
			t.field[move.y][move.x] = t.humanSymbol
			score := t.minmax(deep+1, !maximize)
			best = min(best, score)
			t.field[move.y][move.x] = EMPTY
		}
	}
	return best
}

func (t *TicTacToe) isComputerWon() bool {
	return isWinner(t.field, t.computerSymbol)
}

func (t *TicTacToe) isHumanWon() bool {
	return isWinner(t.field, t.humanSymbol)
}

func isWinner(f [3][3]rune, symbol rune) bool {
	for y := range 3 {
		if f[y][0] == f[y][1] && f[y][1] == f[y][2] && f[y][0] == symbol {
			return true
		}
	}
	for x := range 3 {
		if f[0][x] == f[1][x] && f[1][x] == f[2][x] && f[0][x] == symbol {
			return true
		}
	}
	if f[0][0] == f[1][1] && f[1][1] == f[2][2] && f[0][0] == symbol {
		return true
	}
	if f[0][2] == f[1][1] && f[1][1] == f[2][0] && f[0][2] == symbol {
		return true
	}
	return false
}

func (t *TicTacToe) avaliableMoves() []move {
	moves := make([]move, 0)
	for y := range 3 {
		for x := range 3 {
			if t.field[y][x] == EMPTY {
				moves = append(moves, move{y: y, x: x})
			}
		}
	}
	return moves
}

func (t *TicTacToe) String() string {
	var sb strings.Builder
	for y := range 3 {
		for x := range 3 {
			sb.WriteRune(t.field[y][x])
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}
