package main

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/lipgloss"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	backgroundColor = "#000000"
	landColor       = "#4C956C"
	enemyColor      = "#00FF0F"
	waterColor      = "#0B598D"
	wallColor       = "#A89D9E"
	outboundColor   = "#FF0090"
	attackPointer   = "#FF0000"
	playerColor     = "#FFFF00"
	borderColor     = "#322F20"
	attackColor     = "#7B0828"
	letterColor     = "#322F20"
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
	//fmt.Println("Start")
	mB := NewModelBattlefield()

	//fmt.Println("MB DONE")
	mO := NewModelOptions()

	//fmt.Println("MO DONE")
	mL := NewModelLogger()

	//fmt.Println("ML DONE")
	m := &model{
		Battlefield: mB,
		OptionsList: mO,
		Logger:      mL,
		State:       BATTLEFIELD,
	}
	m.Battlefield.Parent = m

	m.OptionsList.Parent = m

	mB.applyColorChangeInit()
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
		currentPosition := m.Battlefield.Character.Position
		m.Battlefield = *GetBattlefieldType(m.Battlefield.Update(msg))
		if currentPosition != m.Battlefield.Character.Position {
			m.Logger.AddToLog(fmt.Sprintf("Player moved to new position: %v", m.Battlefield.Character.Position))
		}
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
