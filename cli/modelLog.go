package main

import (
	_ "fmt"
	tea "github.com/charmbracelet/bubbletea"
	"strings"
)

const LinesInLog = 7

type modelLog struct {
	Log         []string
	EscapeCodes []string
}

func (m modelLog) Init() tea.Cmd {
	return nil
}

func (m *modelLog) View() string {
	var b strings.Builder
	lenOfLog := len(m.Log)
	for i, c := lenOfLog-1, 0; i >= 0; i, c = i-1, c+1 {
		b.WriteString(m.EscapeCodes[c] + m.Log[i] + "\033[0m\n")
	}
	return b.String()
}

func (m *modelLog) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m *modelLog) AddToLog(msg string) {
	msg = "•" + msg
	if len(m.Log) >= 3 {
		m.Log = append(m.Log[1:], msg)
	} else {
		m.Log = append(m.Log, msg)
	}
}