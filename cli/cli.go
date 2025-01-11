package main

import (
	"PostRPG/Battlefield"
	_ "fmt"
	"log"
	"os"
	"strconv"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	landColor     = "#4C956C"
	waterColor    = "#0B598D"
	wallColor     = "#A89D9E"
	outboundColor = "#FF0090"
	playerColor   = "#FF0000"
	borderColor   = "#322F20"
	letterColor   = "#322F20"
)

type State uint8

const (
	BATTLEFIELD State = iota
	OPTIONS
)

type model struct {
	Battlefield modelBattlefield
	OptionsList modelOptions
	Logger      modelLog
	State       State
}

// TODO:
func DefaultStyles(m model) lipgloss.Style {
	re := lipgloss.NewRenderer(os.Stdout)
	baseStyle := re.NewStyle().Border(lipgloss.DoubleBorder())

	return baseStyle
}

func NewModel() *model {
	m := &model{
		Battlefield: modelBattlefield{
			Bfield: Battlefield.NewBattleField(2, 3),
			Cursor: struct{ x, y int }{x: 0, y: 1},
		},
		OptionsList: modelOptions{
			Options: []string{
				"USE SKILL",
				"USE WEAPON",
			},
			OptionsCursor: 0,
		},
		Logger: modelLog{
			Log: []string{"", "", "", "", "", "", ""},
			EscapeCodes: []string{
				"\033[38;5;255m", // Bright White
				"\033[38;5;252m", // Light Gray
				"\033[38;5;246m", // Gray
				"\033[38;5;240m", // Dark Gray
				"\033[38;5;238m", // Darker Gray
				"\033[38;5;236m", // Very Dark Gray
				"\033[38;5;234m", // Almost Black
			},
		},
		State: BATTLEFIELD,
	}

	// Table Styling
	m.Battlefield.Table = table.New().Border(lipgloss.NormalBorder()).BorderRow(true)
	for rIdx, row := range m.Battlefield.Bfield {
		var newRow []string
		for cIdx := range row {
			i := strconv.Itoa(m.Battlefield.Bfield[rIdx][cIdx])
			newRow = append(newRow, i)
		}
		m.Battlefield.Table.Row(newRow...)
	}
	m.Battlefield.applyColorChange()

	Battlefield.LogBattlefield(m.Battlefield.Bfield)
	return m
}

func (m *model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab":
			if m.State == BATTLEFIELD {
				m.State = OPTIONS
				m.Logger.AddToLog("You are now in Options")
			} else {
				m.State = BATTLEFIELD
				m.Logger.AddToLog("You are now in Battlefield")
			}
		case "q", "Q":
			return m, tea.Quit
		}
	}

	switch m.State {
	case BATTLEFIELD:
		m.Battlefield = *GetBattlefieldType(m.Battlefield.Update(msg))
		return m, nil
	case OPTIONS:
		m.OptionsList = *GetOptionsType(m.OptionsList.Update(msg))
		return m, nil
	}
	return m, nil
}

func (m *model) View() string {
	return lipgloss.JoinHorizontal(
		lipgloss.Center,
		m.Battlefield.View(),
		lipgloss.NewStyle().Border(lipgloss.NormalBorder()).Render(m.OptionsList.View()),
		lipgloss.NewStyle().Border(lipgloss.NormalBorder()).Width(50).Render(m.Logger.View()))
}

func Run() {
	m := NewModel()
	p := tea.NewProgram(m, tea.WithAltScreen())
	_, err := p.Run()
	if err != nil {
		log.Fatalln("(-) Error starting the program: ", err)
	}
}

func main() { Run() }
