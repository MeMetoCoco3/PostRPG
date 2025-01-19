package main

import (
	"PostRPG/Battlefield"
	"PostRPG/internal/database"
	"fmt"
	_ "log"
	"strconv"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"

	tea "github.com/charmbracelet/bubbletea"
)

type modelBattlefield struct {
	Bfield    [][]int
	Character Character
	Enemies   []*Character
	Parent    *model
	//Style      lipgloss.Style
	Table *table.Table
	DB    *database.Queries
}

func (m *modelBattlefield) Init() tea.Cmd {
	return nil
}

func NewModelBattlefield() modelBattlefield {
	dbConexion := GetConexion()
	enemies := GetEnemies(3, []int{5, 6, 7}, []int{5, 6, 7})

	mB := modelBattlefield{
		Bfield:    Battlefield.NewBattleField(2, 3),
		Character: *NewCharacter("Vidal", WARRIOR, "$"),
		Enemies:   enemies,
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

	Battlefield.LogBattlefield(mB.Bfield)

	return mB
}

func (m *modelBattlefield) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// TODO: do not like how i deal with directions here, is different than in modeloptions
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
			for _, enemy := range m.Enemies {
				if enemy.Position.X == nextX && enemy.Position.Y == nextY {
					m.Parent.Logger.AddToLog("Colision against " + enemy.Name)
					return m, nil
				}
			}
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
	// Maybe in the future i will take out everything related to paint enemies and player out of here and put it outside,
	// because its to much checkking

	var colorStyle lipgloss.Style
	m.Table.StyleFunc(func(row, col int) lipgloss.Style {

		if col == m.Character.Position.X && row == m.Character.Position.Y {
			return colorStyle.Background(lipgloss.Color(playerColor)).Padding(0, 1, 0).Bold(true)
		} else {

			for _, enemy := range m.Enemies {
				if col == enemy.Position.X && row == enemy.Position.Y {
					return colorStyle.Background(lipgloss.Color(enemyColor)).Padding(0, 1, 0).Bold(true)
				}
			}
		}

		if m.Parent.OptionsList.OptionsCursor == 1 && m.Parent.State == OPTIONS {
			for _, position := range *m.Parent.OptionsList.AttackMode {
				if col == position.X && row == position.Y {

					m.Parent.Logger.AddToLog(fmt.Sprintf("P:%v%v, CR: %v%v", position.X, position.Y, col, row))
					return colorStyle.Background(lipgloss.Color(attackColor)).Padding(0, 1, 0).Bold(true)
				}
			}
		}

		switch {
		case m.Bfield[row-1][col] == 0:
			return colorStyle.Foreground(lipgloss.Color(landColor)).Padding(0, 1, 0).Bold(true).Background(lipgloss.Color(backgroundColor))
		case m.Bfield[row-1][col] == 1:
			return colorStyle.Foreground(lipgloss.Color(waterColor)).Padding(0, 1, 0).Bold(true).Background(lipgloss.Color(backgroundColor))
		case m.Bfield[row-1][col] == 2:
			return colorStyle.Foreground(lipgloss.Color(wallColor)).Padding(0, 1, 0).Bold(true).Background(lipgloss.Color(backgroundColor))
		default:
			return colorStyle.Foreground(lipgloss.Color(outboundColor)).Padding(0, 1, 0).Bold(true).Background(lipgloss.Color(backgroundColor))
		}
	})

}

func (m *modelBattlefield) applyColorChangeInit() {
	var colorStyle lipgloss.Style
	m.Table.StyleFunc(func(row, col int) lipgloss.Style {

		if col == m.Character.Position.X && row == m.Character.Position.Y {
			return colorStyle.Background(lipgloss.Color(playerColor)).Padding(0, 1, 0).Bold(true)
		} else {

			for _, enemy := range m.Enemies {
				if col == enemy.Position.X && row == enemy.Position.Y {
					return colorStyle.Background(lipgloss.Color(enemyColor)).Padding(0, 1, 0).Bold(true)
				}
			}
		}

		switch {
		case m.Bfield[row-1][col] == 0:
			return colorStyle.Foreground(lipgloss.Color(landColor)).Padding(0, 1, 0).Bold(true).Background(lipgloss.Color(backgroundColor))
		case m.Bfield[row-1][col] == 1:
			return colorStyle.Foreground(lipgloss.Color(waterColor)).Padding(0, 1, 0).Bold(true).Background(lipgloss.Color(backgroundColor))
		case m.Bfield[row-1][col] == 2:
			return colorStyle.Foreground(lipgloss.Color(wallColor)).Padding(0, 1, 0).Bold(true).Background(lipgloss.Color(backgroundColor))
		default:
			return colorStyle.Foreground(lipgloss.Color(outboundColor)).Padding(0, 1, 0).Bold(true).Background(lipgloss.Color(backgroundColor))
		}
	})

}

func (m *modelBattlefield) DeleteEnemy(index int) {
	enemies := m.Enemies
	if index == len(enemies)-1 {
		enemies = enemies[:index]
	} else {
		enemies = append(enemies[:index], enemies[index+1:]...)
	}

	m.Enemies = enemies
}
