package test

import (
	"testing"
	"time"

	tuitest "github.com/aschey/tui-tester"
	"github.com/stretchr/testify/require"
)

func TestSendInputAndExpectOutput(t *testing.T) {
	tester, err := tuitest.NewTester("./testapp", tuitest.WithErrorHandler(func(err error) error {
		t.Error(err)
		return err
	}))

	if err != nil {
		t.Error(err)
	}
	console, _ := tester.CreateConsole()
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

func TestMinInputInterval(t *testing.T) {
	tester, err := tuitest.NewTester("./testapp", tuitest.WithMinInputInterval(100*time.Millisecond))
	if err != nil {
		t.Error(err)
	}
	console, _ := tester.CreateConsole()
	start := time.Now()
	console.SendString("a")
	console.SendString("b")
	end := time.Now()
	duration := end.Sub(start)
	require.True(t, duration >= 100*time.Millisecond)
	require.True(t, duration < 200*time.Millisecond)
}
