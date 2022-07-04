package test

import (
	"testing"

	tuitest "github.com/aschey/tui-tester"
)

func TestTester(t *testing.T) {
	tester, err := tuitest.NewTester("./testapp", tuitest.WithErrorHandler(func(err error) error {
		t.Error(err)
		return err
	}))

	if err != nil {
		t.Error(err)
	}
	console, _ := tester.CreateConsole([]string{})
	console.TrimOutput = true

	// Wait for initialization
	_, _ = console.WaitFor(func(state tuitest.TermState) bool {
		return state.Output() == "You typed:"
	})

	console.SendString("input")
	_, _ = console.WaitFor(func(state tuitest.TermState) bool {
		return state.Output() == "You typed: input"
	})
	if err != nil {
		t.Error(err)
	}

	console.SendString(tuitest.KeyCtrlC)
	_ = console.WaitForTermination()

	_ = tester.TearDown()
}
