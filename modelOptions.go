package main

import (
	"PostRPG/Battlefield"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type modelOptions struct {
	Options       []string
	OptionsCursor int
	OnAttackMode  bool
	AttackMode    *[]Position
	Parent        *model
}

func NewModelOptions() modelOptions {
	return modelOptions{
		Options: []string{
			"USE SKILL",
			"USE WEAPON",
			"SAVE",
		},
		OptionsCursor: 0,
		OnAttackMode:  false,
	}
}

func (m modelOptions) Init() tea.Cmd {
	return nil
}

func (m *modelOptions) View() string {
	var b strings.Builder
	for i, val := range m.Options {
		if i == m.OptionsCursor {
			b.WriteString("> " + val + "\n")
		} else {
			b.WriteString("  " + val + "\n")
		}
	}
	return b.String()
}

func (m *modelOptions) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.OptionsCursor <= 0 {
				m.OptionsCursor = 0
				return m, nil
			}
			m.OptionsCursor--
		case "down", "j":
			if m.OptionsCursor >= len(m.Options)-1 {
				m.OptionsCursor = len(m.Options) - 1
				return m, nil
			}
			m.OptionsCursor++
		case "enter":
			switch m.Options[m.OptionsCursor] {
			case "USE SKILL":
				m.Parent.Logger.AddToLog("We are using a skill.")
			case "USE WEAPON":
				x := m.Parent.Battlefield.Character.Position.X
				y := m.Parent.Battlefield.Character.Position.Y
				m.OnAttackMode = true

				// CHECK direction of ATTACK
				cnt := 0
				for _, direction := range Directions {
					checkPosX := x + direction.X
					checkPosY := y + direction.Y
					if val := Battlefield.CheckNextPosition(m.Parent.Battlefield.Bfield, checkPosX, checkPosY); val != 3 {
						(*m.AttackMode)[cnt] = Position{X: checkPosX, Y: checkPosX}
					}
				}

				/*
					for i, enemy := range m.Parent.Battlefield.Enemies {
						if dx, dy := DistanceBetweenTwoPoints(x, y, enemy.Position.X, enemy.Position.Y); (dx == 0 && dy == 1) || (dy == 0 && dx == 1) {
							m.Parent.Logger.AddToLog("We are attaking " + enemy.Name)
							m.Parent.Battlefield.DeleteEnemy(i)
						}
					}
				*/
			case "SAVE":
				m.Parent.Logger.AddToLog("We are saving.")
			}
		}
	}
	return m, nil
}

func (m *modelOptions) DeleteAttackMode() {

	m.OnAttackMode = false
	emptyAttack := make([]Position, 4)
	m.Parent.Battlefield.AttackMode = &emptyAttack
	m.AttackMode = &emptyAttack
}

func GetOptionsType(m tea.Model, c tea.Cmd) *modelOptions {
	optionsModel := m.(*modelOptions)
	return optionsModel
}
