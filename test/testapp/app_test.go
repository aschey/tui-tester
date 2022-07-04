package testapp

import (
	"fmt"
	"os"
	"testing"

	tuitest "github.com/aschey/tui-tester"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	input string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyRunes, tea.KeySpace:
			m.input += msg.String()
		}
	}
	return m, nil
}

func (m model) View() string {
	return "You typed: " + m.input
}

func TestApp(t *testing.T) {
	_, _ = tuitest.NewTester(".")

	if err := tea.NewProgram(model{input: ""}).Start(); err != nil {
		fmt.Printf("Could not start program :(\n%v\n", err)
		os.Exit(1)
	}
}
