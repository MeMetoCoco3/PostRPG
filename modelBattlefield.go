package main

import (
	"PostRPG/Battlefield"
	"PostRPG/internal/database"
	_ "fmt"
	"strconv"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"

	tea "github.com/charmbracelet/bubbletea"
)

type modelBattlefield struct {
	Bfield    [][]int
	Character Character
	Enemies   []struct {
		x int
		y int
	}
	Style lipgloss.Style
	Table *table.Table
	DB    *database.Queries
}

func (m *modelBattlefield) Init() tea.Cmd {
	return nil
}

func NewModelBattlefield() modelBattlefield {
	dbConexion := GetConexion()

	mB := modelBattlefield{
		Bfield:    Battlefield.NewBattleField(2, 3),
		Character: *NewCharacter("Vidal", WARRIOR, "$"),
		Enemies:   []struct{ x, y int }{{x: 5, y: 7}, {x: 4, y: 7}, {x: 6, y: 7}, {x: 5, y: 8}},
		DB:        dbConexion,
	}

	mB.Character.Position = Position{0, 1}

	// Table Styling
	mB.Table = table.New().Border(lipgloss.NormalBorder()).BorderRow(true)
	for rIdx, row := range mB.Bfield {
		var newRow []string
		for cIdx := range row {
			i := strconv.Itoa(mB.Bfield[rIdx][cIdx])
			newRow = append(newRow, i)
		}
		mB.Table.Row(newRow...)
	}
	mB.applyColorChange()

	Battlefield.LogBattlefield(mB.Bfield)

	return mB
}

func (m *modelBattlefield) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	directions := [][]int{
		{0, -1},
		{1, 0},
		{0, 1},
		{-1, 0},
	}
	var nextX, nextY int
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			nextX, nextY = m.Character.Position.X+directions[0][0], m.Character.Position.Y+directions[0][1]
		case "down", "j":
			nextX, nextY = m.Character.Position.X+directions[2][0], m.Character.Position.Y+directions[2][1]
		case "left", "a":
			nextX, nextY = m.Character.Position.X+directions[3][0], m.Character.Position.Y+directions[3][1]
		case "right", "d":
			nextX, nextY = m.Character.Position.X+directions[1][0], m.Character.Position.Y+directions[1][1]
		}
		nextPosition := Battlefield.CheckNextPosition(m.Bfield, nextX, nextY-1)

		if nextPosition == Battlefield.LAND {
			m.Character.Position.X = nextX
			m.Character.Position.Y = nextY
			m.applyColorChange()

		}
		return m, nil
	}
	return m, nil
}

func GetBattlefieldType(m tea.Model, c tea.Cmd) *modelBattlefield {
	battlefieldModel := m.(*modelBattlefield)
	return battlefieldModel
}

func (m *modelBattlefield) View() string {
	return m.Table.Render()
}

func (m *modelBattlefield) applyColorChange() {
	m.Table.StyleFunc(func(row, col int) lipgloss.Style {
		var colorStyle lipgloss.Style
		if col == m.Character.Position.X && row == m.Character.Position.Y {
			return colorStyle.Background(lipgloss.Color(playerColor)).Padding(0, 1, 0).Bold(true)
		} else {
			for _, enemy := range m.Enemies {
				if col == enemy.x && row == enemy.y {
					return colorStyle.Background(lipgloss.Color(enemyColor)).Padding(0, 1, 0).Bold(true)
				}
			}
		}

		switch {
		case m.Bfield[row-1][col] == 0:
			return colorStyle.Foreground(lipgloss.Color(landColor)).Padding(0, 1, 0).Bold(true)
		case m.Bfield[row-1][col] == 1:
			return colorStyle.Foreground(lipgloss.Color(waterColor)).Padding(0, 1, 0).Bold(true)
		case m.Bfield[row-1][col] == 2:
			return colorStyle.Foreground(lipgloss.Color(wallColor)).Padding(0, 1, 0).Bold(true)
		default:
			return colorStyle.Foreground(lipgloss.Color(outboundColor)).Padding(0, 1, 0).Bold(true)
		}
	})
}
