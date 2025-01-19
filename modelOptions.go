package main

import (
	"PostRPG/Battlefield"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"strings"
)

type modelOptions struct {
	Options        []string
	OptionsCursor  int
	AimCursor      int
	Aiming         bool
	AimPosition    *Character
	EnemiesOnRange *[]Character
	AttackMode     *[]Position
	Parent         *model
}

func NewModelOptions() modelOptions {
	mO := modelOptions{
		Options: []string{
			"USE SKILL",
			"USE WEAPON",
			"SAVE",
		},
		OptionsCursor: 0,
		AimCursor:     0,
		Aiming:        false,
	}
	mO.DeleteAttackMode()
	return mO
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
			if m.Aiming {
				m.DeleteAttackMode()
				return m, nil
			}
			return m, tea.Quit
		case "up", "k":
			if m.Aiming {
				return m, nil
			}
			if m.OptionsCursor <= 0 {
				m.OptionsCursor = 0
				return m, nil
			}
			m.OptionsCursor--
		case "down", "j":
			if m.Aiming {
				return m, nil
			}
			if m.OptionsCursor >= len(m.Options)-1 {
				m.OptionsCursor = len(m.Options) - 1
				return m, nil
			}
			m.OptionsCursor++

		case "left", "a":
			closeEnemiesCounter := len(*m.EnemiesOnRange)
			if m.Aiming && closeEnemiesCounter > 0 {
				if m.AimCursor <= 0 {
					m.AimCursor = closeEnemiesCounter - 1
				} else {
					m.AimCursor--
				}
				m.AimPosition = &(*m.EnemiesOnRange)[m.AimCursor]
				return m, nil
			}

		case "right", "d":
			closeEnemiesCounter := len(*m.EnemiesOnRange)
			if m.Aiming && closeEnemiesCounter > 0 {
				if m.AimCursor >= closeEnemiesCounter-1 {
					m.AimCursor = 0
				} else {
					m.AimCursor++
				}
				m.AimPosition = &(*m.EnemiesOnRange)[m.AimCursor]
				return m, nil
			}

		case "enter":
			if m.Aiming {
				return m, nil
			}
			switch m.Options[m.OptionsCursor] {
			case "USE SKILL":
				m.Parent.Logger.AddToLog("We are using a skill.")
			case "USE WEAPON":
				x := m.Parent.Battlefield.Character.Position.X
				y := m.Parent.Battlefield.Character.Position.Y

				// CHECK direction of ATTACK
				for _, direction := range Directions {

					checkPosX := x + direction.X
					checkPosY := y + direction.Y
					if val := Battlefield.CheckNextPosition(m.Parent.Battlefield.Bfield, checkPosX, checkPosY); val != 3 || checkPosY == LenBattlefield {
						(*m.AttackMode) = append((*m.AttackMode), Position{X: checkPosX, Y: checkPosY})
						for _, enemy := range m.Parent.Battlefield.Enemies {
							if enemy.Position.X == checkPosX && enemy.Position.Y == checkPosY {
								(*m.EnemiesOnRange) = append((*m.EnemiesOnRange), *enemy)
								m.AimPosition = &(*m.EnemiesOnRange)[0]
							}
						}
					}
					m.Aiming = true
				}
			case "SAVE":
				m.Parent.Logger.AddToLog("We are saving.")
			}
		}
	}
	return m, nil
}

func (m *modelOptions) DeleteAttackMode() {
	emptyAttack := []Position{}
	emptyEnemiesOnRange := []Character{}
	m.Aiming = false
	m.AttackMode = &emptyAttack
	m.EnemiesOnRange = &emptyEnemiesOnRange
}

func GetOptionsType(m tea.Model, c tea.Cmd) *modelOptions {
	optionsModel := m.(*modelOptions)
	return optionsModel
}

func (m *modelOptions) GetEnemiesOnSight() []Character {

	fmt.Println("Jamones")
	return nil
}
