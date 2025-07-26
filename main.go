package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	Banner = ` 


████████╗███████╗██╗  ██╗████████╗██████╗  ██████╗ ██╗  ██╗
╚══██╔══╝██╔════╝╚██╗██╔╝╚══██╔══╝██╔══██╗██╔═══██╗╚██╗██╔╝
   ██║   █████╗   ╚███╔╝    ██║   ██████╔╝██║   ██║ ╚███╔╝ 
   ██║   ██╔══╝   ██╔██╗    ██║   ██╔══██╗██║   ██║ ██╔██╗ 
   ██║   ███████╗██╔╝ ██╗   ██║   ██████╔╝╚██████╔╝██╔╝ ██╗
   ╚═╝   ╚══════╝╚═╝  ╚═╝   ╚═╝   ╚═════╝  ╚═════╝ ╚═╝  ╚═╝
                                                           
  `
	purple = lipgloss.Color("99")
	darkBG = lipgloss.Color("#282828")
	style  = lipgloss.NewStyle().
		Bold(true).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(purple).
		Foreground(purple).
		PaddingTop(2).
		PaddingLeft(3).
		Width(90).
		Height(31).
		Align(lipgloss.Center)
	inputStyle = lipgloss.NewStyle().MarginTop(1).Width(40).
			Border(lipgloss.RoundedBorder())
)

type Model struct {
	textarea textarea.Model
}

func initialModel(w, h int) Model {
	ta := textarea.New()
	ta.Placeholder = "Enter"
	ta.Focus()

	ta.Prompt = " => "
	ta.CharLimit = 200
	ta.SetWidth(w)
	ta.SetHeight(h)
	ta.ShowLineNumbers = false

	return Model{textarea: ta}
}

func (m Model) Init() tea.Cmd {
	return textarea.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	m.textarea, cmd = m.textarea.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	view := Banner + "\n" + inputStyle.Render(m.textarea.View())
	return style.Render(view)
}

func main() {
	m := initialModel(5, 1)
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
