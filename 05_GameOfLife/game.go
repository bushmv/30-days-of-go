package main

import "math/rand/v2"

const (
	ALIVE = true
	DEAD  = false
)

type Game struct {
	cur        [][]bool
	next       [][]bool
	Rows       int
	Cols       int
	Generation int
}

func NewGame(rows int, cols int) *Game {
	cur := make([][]bool, 0)
	next := make([][]bool, 0)
	for range rows {
		cur = append(cur, make([]bool, cols))
		next = append(next, make([]bool, cols))
	}
	return &Game{
		cur:        cur,
		next:       next,
		Rows:       rows,
		Cols:       cols,
		Generation: 0,
	}
}

func (g *Game) Update() {
	for y := range g.Rows {
		for x := range g.Cols {
			neighbors := g.countNeighbors(y, x)
			switch {
			case neighbors < 2 || neighbors > 3:
				g.next[y][x] = DEAD
			case g.cur[y][x] == DEAD && neighbors == 3:
				g.next[y][x] = ALIVE
			default:
				g.next[y][x] = g.cur[y][x]
			}
		}
	}
	g.cur, g.next = g.next, g.cur
	g.Generation++
}

func (g *Game) countNeighbors(y, x int) int {
	neighbors := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			cy := (y + dy + g.Rows) % g.Rows
			cx := (x + dx + g.Cols) % g.Cols
			if g.cur[cy][cx] == ALIVE {
				neighbors++
			}
		}
	}
	if g.cur[y][x] == ALIVE {
		neighbors--
	}
	return neighbors
}

func (g *Game) Randomize(chaneToBeAlive float64) {
	for y := range g.Rows {
		for x := range g.Cols {
			if rand.Float64() < chaneToBeAlive {
				g.cur[y][x] = ALIVE
			}
		}
	}
}

func (g *Game) IsStable() bool {
	for y := range g.Rows {
		for x := range g.Cols {
			if g.cur[y][x] != g.next[y][x] {
				return false
			}
		}
	}
	return true
}
