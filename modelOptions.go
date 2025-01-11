package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type modelOptions struct {
	Options       []string
	OptionsCursor int
}

func NewModelOptions() modelOptions {
	return modelOptions{
		Options: []string{
			"USE SKILL",
			"USE WEAPON",
			"SAVE",
		},
		OptionsCursor: 0,
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
		}
	}
	return m, nil
}

func GetOptionsType(m tea.Model, c tea.Cmd) *modelOptions {
	optionsModel := m.(*modelOptions)
	return optionsModel
}
