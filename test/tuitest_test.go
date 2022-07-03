package test

import (
	"testing"

	tuitest "github.com/aschey/tui-tester"
)

func TestTester(t *testing.T) {
	tester, err := tuitest.NewTester("./testapp")
	if err != nil {
		t.Error(err)
	}
	console, err := tester.CreateConsole([]string{})
	console.TrimOutput = true
	if err != nil {
		t.Error(err)
	}

	// Wait for initialization
	_, err = console.WaitFor(func(state tuitest.TermState) bool {
		return state.Output() == "You typed:"
	})
	if err != nil {
		t.Error(err)
	}

	console.SendString("input")
	_, err = console.WaitFor(func(state tuitest.TermState) bool {
		return state.Output() == "You typed: input"
	})
	if err != nil {
		t.Error(err)
	}

	console.SendString(tuitest.KeyCtrlC)
	err = console.WaitForTermination()
	if err != nil {
		t.Error(err)
	}

	err = tester.TearDown()
	if err != nil {
		t.Error(err)
	}
}
