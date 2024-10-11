package main

import (
	"PostRPG/Battlefield"
	_ "fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"log"
	"os"
	"strconv"
	_ "strings"
)

const (
	landColor     = "#4C956C"
	waterColor    = "#0B598D"
	wallColor     = "#A89D9E"
	outboundColor = "#FF0000"
	borderColor   = "#322F20"
	letterColor   = "#322F20"
)

type model struct {
	bfield [][]int
	cursor struct {
		x int
		y int
	}
	style lipgloss.Style
	table *table.Table
}

/*
	type Styles struct {
		BorderColor lipglosrs.Color
		InputField  lipgloss.Style
	}
*/
//For the future
func DefaultStyles(m model) lipgloss.Style {
	re := lipgloss.NewRenderer(os.Stdout)
	baseStyle := re.NewStyle().Border(lipgloss.DoubleBorder())

	return baseStyle
}

func NewModel() model {
	b := Battlefield.NewBattleField(2, 3)
	m := model{
		bfield: b,
	}
	m.cursor.x = 0
	m.cursor.y = 0
	//m.style = DefaultStyles(m)

	// Table Styling
	m.table = table.New().Border(lipgloss.NormalBorder()).BorderRow(true)
	for rIdx, row := range b {
		var newRow []string
		for cIdx := range row {
			i := strconv.Itoa(b[rIdx][cIdx])
			newRow = append(newRow, i)
		}
		m.table.Row(newRow...)
	}

	m.table.StyleFunc(func(row, col int) lipgloss.Style {
		var colorStyle lipgloss.Style
		switch {

		case m.bfield[row-1][col] == 0:
			return colorStyle.Background(lipgloss.Color(landColor)).Padding(0, 1, 0).Foreground(lipgloss.Color(letterColor))
		case m.bfield[row-1][col] == 1:
			return colorStyle.Background(lipgloss.Color(waterColor)).Padding(0, 1, 0).Foreground(lipgloss.Color(letterColor))
		case m.bfield[row-1][col] == 2:
			return colorStyle.Background(lipgloss.Color(wallColor)).Padding(0, 1, 0).Foreground(lipgloss.Color(letterColor))
		default:
			return colorStyle.Background(lipgloss.Color(outboundColor)).Padding(0, 1, 0).Foreground(lipgloss.Color(letterColor))
		}
	})
	Battlefield.LogBattlefield(m.bfield)
	return m
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	directions := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	var nextX, nextY int
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			nextX, nextY = m.cursor.x+directions[0][0], m.cursor.y+directions[0][1]
		case "down", "j":
			nextX, nextY = m.cursor.x+directions[2][0], m.cursor.y+directions[2][1]

		case "left", "a":
			nextX, nextY = m.cursor.x+directions[3][0], m.cursor.y+directions[3][1]

		case "right", "d":
			nextX, nextY = m.cursor.x+directions[1][0], m.cursor.y+directions[1][1]
		}
		nextPosition := Battlefield.CheckNextPosition(m.bfield, nextX, nextY)
		if nextPosition == Battlefield.LAND {
			m.cursor.x = nextX
			m.cursor.y = nextY
		}
		return m, nil
	}
	return m, nil
}

func (m model) View() string {
	return lipgloss.JoinVertical(
		lipgloss.Center,
		m.table.Render(),
	)
}

func Run() {
	m := NewModel()
	p := tea.NewProgram(&m, tea.WithAltScreen())
	_, err := p.Run()
	if err != nil {
		log.Fatalln("(-) Error starting the program: ", err)
	}
}

func main() {
	Run()
}
