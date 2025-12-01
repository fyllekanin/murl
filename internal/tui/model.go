package tui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/fyllekanin/murl/internal/softwares"
)

type appState int

const (
	choosingSoftware appState = iota
	choosingOperation
	showingList
)

type AppModel struct {
	state           appState
	selectionModel  *SelectionModel
	listModel       *ListModel
	spinner         spinner.Model
	Software        softwares.Software
	Operation       string
	IsQuitting      bool
	softwareChoices []softwares.Software
	err             error
	width           int
	height          int
}

type versionsLoadedMsg struct {
	versions []softwares.SoftwareVersion
}
type errorMsg struct {
	err error
}

func NewAppModel() *AppModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	m := &AppModel{
		softwareChoices: []softwares.Software{
			softwares.NewGoLang(),
			softwares.NewNodeJs(),
		},
		spinner: s,
	}
	m.toSoftwareSelection()
	return m
}

func (m *AppModel) toSoftwareSelection() {
	var items []item
	for _, choice := range m.softwareChoices {
		items = append(items, NewItem(choice.Name(), choice.Name()))
	}
	m.state = choosingSoftware
	m.selectionModel = NewSelectionModel("Which software to work with?", items)
}

func (m *AppModel) toOperationSelection() {
	m.state = choosingOperation
	m.selectionModel = NewSelectionModel("Which operation?", []item{
		NewItem("List available", "list"),
		NewItem("Install", "install"),
	})
}

func (m *AppModel) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m *AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.err != nil {
		if _, ok := msg.(tea.KeyMsg); ok {
			m.err = nil
		}
		return m, nil
	}
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			m.IsQuitting = true
			return m, tea.Quit
		}
	case versionsLoadedMsg:
		m.state = showingList
		var columns []table.Column
		var rows []table.Row
		if m.Software.Name() == "nodejs" {
			columns = []table.Column{
				{Title: "Version", Width: 40},
				{Title: "LTS", Width: 10},
			}
			for _, v := range msg.versions {
				rows = append(rows, table.Row{v.Name, fmt.Sprintf("%t", v.IsLts)})
			}
		} else {
			columns = []table.Column{
				{Title: "Version", Width: 40},
			}
			for _, v := range msg.versions {
				rows = append(rows, table.Row{v.Name})
			}
		}
		m.listModel = NewListModel(columns, rows)
		return m, nil
	case errorMsg:
		m.err = msg.err
		return m, nil
	}

	var cmd tea.Cmd
	switch m.state {
	case choosingSoftware:
		var model tea.Model
		model, cmd = m.selectionModel.Update(msg)
		m.selectionModel = model.(*SelectionModel)

		if cmd != nil && m.selectionModel.IsQuitting {
			m.IsQuitting = true
			return m, tea.Quit
		}
		if cmd != nil {
			sel := m.selectionModel.GetSelection()
			for _, choice := range m.softwareChoices {
				if choice.Name() == sel.Value() {
					m.Software = choice
					break
				}
			}
			m.toOperationSelection()
			return m, nil
		}

	case choosingOperation:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			if msg.String() == "b" {
				m.toSoftwareSelection()
				return m, nil
			}
		}

		var model tea.Model
		model, cmd = m.selectionModel.Update(msg)
		m.selectionModel = model.(*SelectionModel)
		if cmd != nil && m.selectionModel.IsQuitting {
			m.IsQuitting = true
			return m, tea.Quit
		}

		if cmd != nil {
			m.Operation = m.selectionModel.GetSelection().Value()
			if m.Operation == "list" {
				m.state = showingList // show loading spinner
				return m, func() tea.Msg {
					versions, err := m.Software.List()
					if err != nil {
						return errorMsg{err: err}
					}
					return versionsLoadedMsg{versions: versions}
				}
			}
			m.IsQuitting = true
			return m, tea.Quit
		}
	case showingList:
		if m.listModel == nil { // Still loading
			m.spinner, cmd = m.spinner.Update(msg)
			return m, cmd
		}
		switch msg := msg.(type) {
		case tea.KeyMsg:
			if msg.String() == "b" || msg.String() == "q" {
				m.toOperationSelection()
				return m, nil
			}
		}
		var model tea.Model
		model, cmd = m.listModel.Update(msg)
		m.listModel = model.(*ListModel)
		return m, cmd
	}

	return m, nil
}

func (m *AppModel) footerView() string {
	footerStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	var help string
	switch m.state {
	case choosingSoftware:
		help = "up/down: navigate, enter: select, ctrl+c: quit"
	case choosingOperation:
		help = "up/down: navigate, enter: select, b: back, ctrl+c: quit"
	case showingList:
		if m.listModel == nil {
			help = "ctrl+c: quit"
		} else {
			help = "up/down: navigate, q/b: back, ctrl+c: quit"
		}
	}
	return footerStyle.Render(help)
}

func (m *AppModel) View() string {
	if m.IsQuitting {
		return "Quitting..."
	}
	if m.err != nil {
		errorStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("9"))
		return errorStyle.Render("Error: "+m.err.Error()) + "\n\nPress any key to continue."
	}

	var view string
	switch m.state {
	case choosingSoftware:
		view = m.selectionModel.View()
	case choosingOperation:
		view = m.selectionModel.View()
	case showingList:
		if m.listModel == nil {
			view = m.spinner.View() + " Loading versions..."
		} else {
			view = m.listModel.View()
		}
	}

	return lipgloss.JoinVertical(lipgloss.Left,
		view,
		m.footerView(),
	)
}
