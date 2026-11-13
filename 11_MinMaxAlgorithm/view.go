package main

import (
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	game *TicTacToe
	y    int
	x    int
	msg  string
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
		case "up":
			m.y = max(0, m.y-1)
		case "down":
			m.y = min(2, m.y+1)
		case "left":
			m.x = max(0, m.x-1)
		case "right":
			m.x = min(2, m.x+1)
		case "enter":
			err := m.game.move(m.y, m.x)
			if err != nil {
				m.msg = err.Error()
			}
			if m.game.state != stateRunning {
				switch m.game.state {
				case stateWon:
					m.msg = "you win (q to exit)"
				case stateLose:
					m.msg = "you lose (q to exit)"
				case stateDraw:
					m.msg = "draw(q to exit)"
				}
			}
		}
	}
	return m, nil
}

func (m model) View() tea.View {
	f := m.game.field
	var sb strings.Builder

	highlight := lipgloss.NewStyle().Background(lipgloss.Color("#AAA")).Foreground(lipgloss.Color("#1b1b1b"))
	for y := range 3 {
		for x := range 3 {
			if m.y == y && m.x == x {
				sb.WriteString(highlight.Render(string(f[y][x])))
			} else {
				sb.WriteRune(f[y][x])
			}
		}
		sb.WriteRune('\n')
	}
	sb.WriteString(m.msg)
	sb.WriteString("\n")
	return tea.View{
		Content: sb.String(),
	}
}
