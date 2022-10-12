package tuitest

import (
	"fmt"
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

type Color uint16

func (c Color) Int() uint16 {
	return uint16(c)
}

func (c Color) String() string {
	return fmt.Sprint(c.Int())
}

func (t TermState) ForegroundColor(row int, col int) Color {
	_, fg, _ := t.state.Cell(col, row)
	return Color(fg)
}

func (t TermState) BackgroundColor(row int, col int) Color {
	_, _, bg := t.state.Cell(col, row)
	return Color(bg)
}
