package main

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"

	"PostRPG/Battlefield"

	tea "github.com/charmbracelet/bubbletea"
)

type modelBattlefield struct {
	Bfield [][]int
	Cursor struct {
		x int
		y int
	}
	Style lipgloss.Style
	Table *table.Table
}

func (m *modelBattlefield) Init() tea.Cmd {
	return nil
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
			nextX, nextY = m.Cursor.x+directions[0][0], m.Cursor.y+directions[0][1]
		case "down", "j":
			nextX, nextY = m.Cursor.x+directions[2][0], m.Cursor.y+directions[2][1]
		case "left", "a":
			nextX, nextY = m.Cursor.x+directions[3][0], m.Cursor.y+directions[3][1]

		case "right", "d":
			nextX, nextY = m.Cursor.x+directions[1][0], m.Cursor.y+directions[1][1]
		}
		nextPosition := Battlefield.CheckNextPosition(m.Bfield, nextX, nextY-1)

		fmt.Printf("%d:%d= %d", nextX, nextY, nextPosition)
		if nextPosition == Battlefield.LAND {
			m.Cursor.x = nextX
			m.Cursor.y = nextY
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
		if col == m.Cursor.x && row == m.Cursor.y {
			return colorStyle.Background(lipgloss.Color(playerColor)).Padding(0, 1, 0).Bold(true)
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
