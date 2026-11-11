package main

import (
	"fmt"
	"strings"
	"time"

	tea "charm.land/bubbletea/v2"
)

type Model struct {
	game   *Game
	paused bool
	delay  time.Duration
}

func NewModel(game *Game, paused bool, delay time.Duration) Model {
	return Model{
		game:   game,
		paused: paused,
		delay:  delay,
	}
}

type TickMsg struct{}

func tick(delay time.Duration) tea.Cmd {
	return tea.Tick(delay, func(time.Time) tea.Msg {
		return TickMsg{}
	})
}

func (m Model) Init() tea.Cmd {
	return tick(m.delay)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "space":
			m.paused = !m.paused
			return m, nil
		case "r":
			m.game.Randomize(0.5)
			return m, nil
		case "s":
			if m.paused {
				m.game.Update()
			}
			return m, nil
		case "u":
			m.delay = max(m.delay/2, 100)
			return m, nil
		case "j":
			m.delay = min(m.delay*2, time.Second)
			return m, nil
		}
	case TickMsg:
		if !m.paused {
			m.game.Update()
		}
		return m, tick(m.delay)
	}
	return m, nil
}

func (m Model) View() tea.View {
	status := "RUNNING"
	if m.paused {
		status = "PAUSED"
	}

	info := fmt.Sprintf("Generation: %d, %s\n", m.game.Generation, status)
	help := "space = pause | r - reset | s - step | u - faster | j - slower | q - exit\n\n"
	var field strings.Builder
	for y := 0; y < m.game.Rows; y++ {
		for x := 0; x < m.game.Cols; x++ {
			if m.game.cur[y][x] {
				field.WriteString("#")
			} else {
				field.WriteString(" ")
			}
		}
		field.WriteByte('\n')
	}
	content := info + help + field.String()
	return tea.NewView(content)

}
