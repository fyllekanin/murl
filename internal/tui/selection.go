package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type item struct {
	title string
	value string
}

func NewItem(title, value string) item {
	return item{title: title, value: value}
}

func (i item) Title() string {
	return i.title
}

func (i item) Value() string {
	return i.value
}

type SelectionModel struct {
	title      string
	cursor     int
	items      []item
	IsQuitting bool
}

func NewSelectionModel(title string, items []item) *SelectionModel {
	return &SelectionModel{
		title: title,
		items: items,
	}
}

func (m *SelectionModel) GetSelection() item {
	if m.cursor >= 0 && m.cursor < len(m.items) {
		return m.items[m.cursor]
	}
	return item{}
}

func (m *SelectionModel) Init() tea.Cmd {
	return nil
}

func (m *SelectionModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.IsQuitting = true
			return m, tea.Quit

		case "enter":
			return m, tea.Quit

		case "down", "j":
			m.cursor++
			if m.cursor >= len(m.items) {
				m.cursor = 0
			}

		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(m.items) - 1
			}
		}
	}

	return m, nil
}

func (m *SelectionModel) View() string {
	var sb strings.Builder
	titleStyle := lipgloss.NewStyle().Bold(true)
	sb.WriteString(titleStyle.Render(m.title) + "\n\n")

	for i, item := range m.items {
		selected := i == m.cursor
		checked := "[ ]"
		if selected {
			checked = lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Render("[x]")
		}
		sb.WriteString(fmt.Sprintf("%s %s\n", checked, item.Title()))
	}

	return sb.String()
}

