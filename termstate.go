package tuitest

import (
	"strings"

	"github.com/ActiveState/vt10x"
)

const DefaultFG = uint16(vt10x.DefaultFG)
const DefaultBG = uint16(vt10x.DefaultBG)

type TermState struct {
	text  string
	state *vt10x.State
}

func (t TermState) Output() string {
	return t.text
}

func (t TermState) OutputLines() []string {
	return strings.Split(t.text, "\n")
}

func (t TermState) NthOutputLine(index int) string {
	outputLines := t.OutputLines()
	if index > len(outputLines)-1 {
		return ""
	}
	return outputLines[index]
}

func (t TermState) NumLines() int {
	return len(t.OutputLines())
}

func (t TermState) FgColor(row int, col int) uint16 {
	_, fg, _ := t.state.Cell(col, row)
	return uint16(fg)
}

func (t TermState) BgColor(row int, col int) uint16 {
	_, _, bg := t.state.Cell(col, row)
	return uint16(bg)
}
