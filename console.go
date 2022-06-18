package tuitest

import (
	"fmt"
	"time"

	"github.com/aschey/termtest"
)

type Console struct {
	console    *termtest.ConsoleProcess
	last       string
	lastInput  time.Time
	Timeout    time.Duration
	TrimOutput bool
	OnError    func(err error) error
}

func (c *Console) wait() {
	remaining := time.Since(c.lastInput)
	if remaining < time.Millisecond {
		time.Sleep(time.Millisecond - remaining)
	}

	c.lastInput = time.Now()
}

func (c *Console) SendBytes(input []byte) {
	c.wait()
	c.console.SendUnterminated(string(input))
}

func (c *Console) SendByte(input byte) {
	c.wait()
	c.console.SendUnterminated(string(rune(input)))
}

func (c *Console) SendString(input string) {
	c.wait()
	c.console.SendUnterminated(input)
}

func (c *Console) WaitFor(condition func(state TermState) bool) (TermState, error) {
	return c.waitFor(condition, nil)
}

func (c *Console) WaitForDuration(condition func(state TermState) bool, duration time.Duration) (TermState, error) {
	return c.waitFor(condition, &duration)
}

func (c *Console) waitFor(condition func(state TermState) bool, duration *time.Duration) (TermState, error) {
	timeout := time.After(c.Timeout)
	if duration != nil && *duration*2 > c.Timeout {
		timeout = time.After(*duration * 2)
	}

	errCh := make(chan error, 1)
	outCh := make(chan TermState, 1)
	go func() {
		_, err := c.console.ExpectCustom(Matcher(condition, c.TrimOutput, outCh, duration))
		if err != nil {
			errCh <- err
		}
	}()

	for {
		select {
		case output := <-outCh:
			return output, nil
		case err := <-errCh:
			return TermState{}, err
		case <-timeout:
			return TermState{}, c.handleError(fmt.Errorf("Timeout exceeded while waiting for condition."))
		}
	}
}

func (c *Console) WaitForTermination() error {
	c.console.Wait(c.Timeout)
	c.console.ExpectExitCode(0)
	c.console.Close()
	return nil
}

func (c *Console) handleError(err error) error {
	if c.OnError != nil && err != nil {
		return c.OnError(err)
	}
	return err
}