package main

import (
	"time"

	tea "charm.land/bubbletea/v2"
)

func main() {
	game := NewGame(20, 20)
	game.Randomize(0.5)
	model := NewModel(game, false, time.Millisecond*500)
	program := tea.NewProgram(model)
	if _, err := program.Run(); err != nil {
		panic(err)
	}
}
