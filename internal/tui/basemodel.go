package tui

import tea "github.com/charmbracelet/bubbletea"

type BaseModel struct {
	cursor     int
	IsQuitting bool
}

func (m BaseModel) Init() tea.Cmd {
	return nil
}

func (m BaseModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			m.IsQuitting = true
			return m, tea.Quit
		case "enter":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m BaseModel) View() string {
	return ""
}
