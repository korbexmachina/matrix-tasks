package main

import (
	"fmt"
	"os"
	// "log"

	utils "github.com/korbexmachina/matrix-tasks/utils"
	tui "github.com/korbexmachina/matrix-tasks/tui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	utils.InitDB("test.db")
	model := tui.InitialModel("test.db")
	p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
