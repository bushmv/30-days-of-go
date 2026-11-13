package main

import (
	tea "charm.land/bubbletea/v2"
)

const (
	stateSearch = iota
	statePathFound
)

type model struct {
	maze         *Maze
	traversal    []point
	path         []point
	traversalIdx int
	state        int
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			switch m.state {
			case stateSearch:
				m.maze.Visit(m.traversal[m.traversalIdx])
				m.traversalIdx++
				if m.traversalIdx == len(m.traversal) {
					m.maze.Clear()
					for _, p := range m.path {
						m.maze.Visit(p)
					}
					m.state = statePathFound
				}
			case statePathFound:

			}
		}
	}
	return m, nil
}

func (m model) View() tea.View {
	return tea.View{
		Content: m.maze.String(),
	}
}
