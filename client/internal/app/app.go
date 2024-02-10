package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dustin-ward/tim/client/internal/style"
)

type AppState int

type AppModel struct {
}

func NewAppModel() AppModel {
	return AppModel{}
}

func (m AppModel) Init() tea.Cmd {
	return nil
}

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		style.WindowWidth = msg.Width
		style.WindowHeight = msg.Height

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			cmds = append(cmds, tea.Quit)
		}
	}

	return m, tea.Batch(cmds...)
}

func (m AppModel) View() string {
	return "Hello"
}
