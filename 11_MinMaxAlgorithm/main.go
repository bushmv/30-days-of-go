package main

import tea "charm.land/bubbletea/v2"

func main() {
	game := NewTicTacToe(true)
	model := model{
		game: game,
		y:    0,
		x:    0,
		msg:  "",
	}
	p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
