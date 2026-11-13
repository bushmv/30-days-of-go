package main

import (
	"container/list"
	"errors"
	"fmt"
	"strings"
)

type pair struct {
	cur  point
	prev point
}

type point struct {
	y int
	x int
}

func (p point) String() string {
	return fmt.Sprintf("y: %d, x: %x", p.y, p.x)
}

const (
	wall = iota
	room
	start
	end
	visited
)

type Maze struct {
	rows      int
	cols      int
	field     [][]int
	start     point
	end       point
	traversal []point
	moves     map[point]point
	path      []point
}

func FromRunes(rows, cols int, data []rune) (*Maze, error) {
	if len(data) != rows*cols {
		return nil, errors.New("no data to construct maze")
	}
	S := point{y: -1, x: -1}
	E := point{y: -1, x: -1}
	field := make([][]int, rows)
	i := 0
	for y := range rows {
		line := make([]int, cols)
		for x := range cols {
			switch r := data[i]; r {
			case '.':
				line[x] = room
			case '#':
				line[x] = wall
			case 'S':
				if S.y != -1 && S.x != -1 {
					return nil, errors.New("more than one start point")
				}
				line[x] = start
				S.y = y
				S.x = x
			case 'E':
				if E.y != -1 && E.x != -1 {
					return nil, errors.New("more than one end point")
				}
				line[x] = end
				E.y = y
				E.x = x
			default:
				return nil, fmt.Errorf("unexpected rune: %c", r)
			}
			i++
		}
		field[y] = line
	}
	return &Maze{
		rows:      rows,
		cols:      cols,
		field:     field,
		start:     S,
		end:       E,
		traversal: make([]point, 0),
		path:      make([]point, 0),
		moves:     make(map[point]point),
	}, nil
}

func (m *Maze) String() string {
	var sb strings.Builder
	for y := range m.rows {
		for x := range m.cols {
			var c rune
			switch m.field[y][x] {
			case wall:
				c = '#'
			case room:
				c = ' '
			case start:
				c = 'S'
			case end:
				c = 'E'
			case visited:
				c = '+'
			}
			sb.WriteRune(c)
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

func (m *Maze) DFS() {
	visited := make(map[point]bool)
	m.rdfs(pair{prev: m.start, cur: m.start}, visited)
	m.restorePath()
}

func (m *Maze) restorePath() {
	cur := m.moves[m.end]
	for cur != m.start {
		m.path = append(m.path, cur)
		cur = m.moves[cur]
	}
	m.field[m.start.y][m.start.x] = start
	m.field[m.end.y][m.end.x] = end
}

func (m *Maze) rdfs(p pair, visited map[point]bool) bool {
	if !m.inside(p.cur) {
		return false
	}
	if m.field[p.cur.y][p.cur.x] == wall {
		return false
	}
	if _, ok := visited[p.cur]; ok {
		return false
	}
	visited[p.cur] = true
	m.traversal = append(m.traversal, p.cur)
	m.moves[p.cur] = p.prev
	if m.field[p.cur.y][p.cur.x] == end {
		return true
	}
	return m.rdfs(pair{prev: p.cur, cur: point{p.cur.y - 1, p.cur.x}}, visited) ||
		m.rdfs(pair{prev: p.cur, cur: point{p.cur.y, p.cur.x + 1}}, visited) ||
		m.rdfs(pair{prev: p.cur, cur: point{p.cur.y + 1, p.cur.x}}, visited) ||
		m.rdfs(pair{prev: p.cur, cur: point{p.cur.y, p.cur.x - 1}}, visited)
}

func (m *Maze) inside(p point) bool {
	return p.x >= 0 && p.x < m.cols && p.y >= 0 && p.y < m.cols
}

func (m *Maze) Clear() {
	for y := range m.rows {
		for x := range m.cols {
			if m.field[y][x] == visited {
				m.field[y][x] = room
			}
		}
	}
	m.field[m.start.y][m.start.x] = start
	m.field[m.end.y][m.end.x] = end
}

func (m *Maze) Visit(p point) {
	m.field[p.y][p.x] = visited
}

func (m *Maze) BFS() {
	q := list.New()
	q.PushBack(pair{m.start, m.start})
	visited := make(map[point]bool)
	for q.Len() > 0 {
		el := q.Front()
		p := q.Remove(el).(pair)
		if !m.inside(p.cur) {
			continue
		}
		if _, ok := visited[p.cur]; ok {
			continue
		}
		if m.field[p.cur.y][p.cur.x] == wall {
			continue
		}
		m.traversal = append(m.traversal, p.cur)
		m.moves[p.cur] = p.prev
		visited[p.cur] = true
		if m.field[p.cur.y][p.cur.x] == end {
			break
		}
		q.PushBack(pair{prev: p.cur, cur: point{p.cur.y - 1, p.cur.x}})
		q.PushBack(pair{prev: p.cur, cur: point{p.cur.y, p.cur.x + 1}})
		q.PushBack(pair{prev: p.cur, cur: point{p.cur.y + 1, p.cur.x}})
		q.PushBack(pair{prev: p.cur, cur: point{p.cur.y, p.cur.x - 1}})
	}
	m.restorePath()
}
