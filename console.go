package tuitest

import (
	"time"

	"github.com/aschey/termtest"
)

type Console struct {
	consoleProcess     *termtest.ConsoleProcess
	last               string
	lastInput          time.Time
	terminationTimeout time.Duration
	TrimOutput         bool
	onError            func(err error) error
	minInputInterval   time.Duration
}

func (c *Console) wait() {
	remaining := time.Since(c.lastInput)
	if remaining < c.minInputInterval {
		time.Sleep(c.minInputInterval - remaining)
	}

	c.lastInput = time.Now()
}

func (c *Console) SendString(input string) {
	c.wait()
	c.consoleProcess.SendUnterminated(input)
}

func (c *Console) WaitFor(condition func(state TermState) bool) (TermState, error) {
	return c.waitFor(condition, nil)
}

func (c *Console) WaitForDuration(condition func(state TermState) bool, duration time.Duration) (TermState, error) {
	return c.waitFor(condition, &duration)
}

func (c *Console) waitFor(condition func(state TermState) bool, duration *time.Duration) (TermState, error) {
	outCh := make(chan TermState, 1)
	_, err := c.consoleProcess.ExpectCustom(Matcher(condition, c.TrimOutput, outCh, duration))
	if err != nil {
		return TermState{}, c.handleError(err)
	}
	return <-outCh, nil
}

func (c *Console) WaitForTermination() error {
	c.consoleProcess.Wait(c.terminationTimeout)
	_, err := c.consoleProcess.ExpectExitCode(0)
	if err != nil {
		return c.onError(err)
	}
	err = c.consoleProcess.Close()
	if err != nil {
		return c.onError(err)
	}
	return nil
}

func (c *Console) handleError(err error) error {
	if c.onError != nil && err != nil {
		return c.onError(err)
	}
	return err
}
