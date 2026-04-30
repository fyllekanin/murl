package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/fyllekanin/murl/internal/tui"
)

func main() {



	
<<<<<<< Updated upstream

	startup := tea.NewProgram(tui.StartUpModel{})

	m, err := startup.Run()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	if m, ok := m.(tui.StartUpModel); ok {
		if m.IsQuitting {
			fmt.Println("Quitting...")
		} else {
			fmt.Println("You selected " + m.GetSelection().Name())
		}
		os.Exit(1)
	}

	/**
	includeUnstable := flag.Bool("unstable", false, "Include unstable versions")
	flag.Parse()
=======
	app := tui.NewAppModel()
	p := tea.NewProgram(app)
>>>>>>> Stashed changes

	m, err := p.Run()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	if m, ok := m.(*tui.AppModel); ok {
		if m.IsQuitting || m.Software == nil || m.Operation == "" {
			fmt.Println("Quitting...")
			os.Exit(0)
		}

		switch m.Operation {
		case "install":
			// Need to get version to install
			fmt.Println("Install not yet implemented in TUI")
		}
	}
	*/
}
<<<<<<< Updated upstream

/**
func getBoolValue(b *bool) bool {
	if b != nil {
		return *b
	}
	return false
}
*/
=======
>>>>>>> Stashed changes
