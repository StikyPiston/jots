package ui

import (
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type pagerModel struct {
	viewport viewport.Model
	content  string
}

func NewPager(content string) pagerModel {
	vp := viewport.New(0, 0)
	vp.SetContent(content)

	return pagerModel{
		viewport: vp,
		content:  content,
	}
}

func (m pagerModel) Init() tea.Cmd {
	return nil
}

func (m pagerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.viewport.Width = msg.Width
		m.viewport.Height = msg.Height
		m.viewport.SetContent(m.content)
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}

	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}

func (m pagerModel) View() string {
	return m.viewport.View() + "\n\n(q to quit)"
}

func RunPager(lines []string) error {
	content := strings.Join(lines, "\n")

	p := tea.NewProgram(
		NewPager(content),
		tea.WithAltScreen(),
	)

	_, err := p.Run()
	return err
}
