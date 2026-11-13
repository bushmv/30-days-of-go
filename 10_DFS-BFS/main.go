package main

import (
	"log"
	"strings"

	tea "charm.land/bubbletea/v2"
)

var s = `
#########
#...S...#
#.#.###.#
#.###..##
#.....###
#..#..E.#
#########`

func main() {
	rows := len(strings.Split(strings.TrimSpace(s), "\n"))
	data := []rune(strings.ReplaceAll(s, "\n", ""))
	cols := len(data) / rows
	maze, err := FromRunes(rows, cols, data)
	if err != nil {
		panic(err)
	}
	maze.BFS()
	// maze.DFS()
	model := model{
		maze:         maze,
		traversal:    maze.traversal,
		path:         maze.path,
		traversalIdx: 0,
		state:        stateSearch,
	}
	p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
