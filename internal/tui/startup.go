package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fyllekanin/murl/internal/softwares"
)

var choices = []softwares.Software{
	softwares.NewGoLang(),
	softwares.NewNodeJs(),
}

type StartUpModel struct {
	BaseModel
}

func (m StartUpModel) GetSelection() softwares.Software {
	return choices[m.cursor]
}

func (m StartUpModel) Init() tea.Cmd {
	return nil
}

func (m StartUpModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	updatedBase, cmd := m.BaseModel.Update(msg)
	m.BaseModel = updatedBase.(BaseModel)
	if cmd != nil {
		return m, cmd
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "down":
			m.cursor++
			if m.cursor >= len(choices) {
				m.cursor = 0
			}

		case "up":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(choices) - 1
			}
		}
	}

	return m, nil
}

func (m StartUpModel) View() string {
	s := "Which software to work with?\n\n"

	for i := 0; i < len(choices); i++ {
		if i == m.cursor {
			s += "[x] "
		} else {
			s += "[ ] "
		}

		s += choices[i].Name() + "\n"
	}

	return s
}
