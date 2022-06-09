package tuitest

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/ActiveState/termtest/expect"
	"github.com/ActiveState/vt10x"
)

type Tester struct {
	doneCh     chan error
	console    *expect.Console
	last       string
	lastInput  time.Time
	Timeout    time.Duration
	TrimOutput bool
	OnError    func(err error) error
}

const DefaultFG = uint16(vt10x.DefaultFG)
const DefaultBG = uint16(vt10x.DefaultBG)

// from https://github.com/acarl005/stripansi/blob/master/stripansi.go
const ansi = "[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))"

var re = regexp.MustCompile(ansi)

func RemoveAnsiSequences(str string) string {
	return re.ReplaceAllString(str, "")
}

func (t *Tester) handleError(err error) error {
	if t.OnError != nil && err != nil {
		return t.OnError(err)
	}
	return err
}

func (t *Tester) wait() {
	remaining := time.Since(t.lastInput)
	if remaining < time.Millisecond {
		time.Sleep(time.Millisecond - remaining)
	}

	t.lastInput = time.Now()
}

func (t *Tester) SendBytes(input []byte) error {
	t.wait()
	_, err := t.console.Send(string(input))
	return t.handleError(err)
}

func (t *Tester) SendByte(input byte) error {
	t.wait()
	_, err := t.console.Send(string(rune(input)))
	return t.handleError(err)
}

func (t *Tester) SendString(input string) error {
	t.wait()
	_, err := t.console.Send(input)
	return t.handleError(err)
}

type outputMatcher struct {
	condition  func(state TermState) bool
	trimOutput bool
	outCh      chan TermState
	duration   *time.Duration
	now        *time.Time
}

func (om *outputMatcher) Match(v interface{}) bool {
	ms, ok := v.(*expect.MatchState)
	if !ok {
		return false
	}

	output := ms.TermState.String()
	if om.trimOutput {
		output = strings.TrimSpace(output)
	}

	state := TermState{text: output, state: ms.TermState}
	result := om.condition(state)
	if om.duration != nil && result {
		if om.now == nil {
			now := time.Now()
			om.now = &now
		}
		result = result && time.Since(*om.now) >= *om.duration
	} else {
		om.now = nil
	}
	if result {
		om.outCh <- state
	}
	return result
}

func (om *outputMatcher) Criteria() interface{} {
	return om.condition
}

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

func (t TermState) FgColor(row int, col int) uint16 {
	_, fg, _ := t.state.Cell(col, row)
	return uint16(fg)
}

func (t TermState) BgColor(row int, col int) uint16 {
	_, _, bg := t.state.Cell(col, row)
	return uint16(bg)
}

func Matcher(condition func(state TermState) bool, trimOutput bool, outCh chan TermState, duration *time.Duration) expect.ExpectOpt {
	return func(opts *expect.ExpectOpts) error {
		opts.Matchers = append(opts.Matchers, &outputMatcher{condition: condition, trimOutput: trimOutput, outCh: outCh, duration: duration})
		return nil
	}
}

func (t *Tester) WaitFor(condition func(state TermState) bool) (TermState, error) {
	return t.waitFor(condition, nil)
}

func (t *Tester) WaitForDuration(condition func(state TermState) bool, duration time.Duration) (TermState, error) {
	return t.waitFor(condition, &duration)
}

func (t *Tester) waitFor(condition func(state TermState) bool, duration *time.Duration) (TermState, error) {
	timeout := time.After(t.Timeout)
	if duration != nil && *duration*2 > t.Timeout {
		timeout = time.After(*duration * 2)
	}

	errCh := make(chan error, 1)
	outCh := make(chan TermState, 1)
	go func() {
		_, err := t.console.Expect(Matcher(condition, t.TrimOutput, outCh, duration))
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
			return TermState{}, t.handleError(fmt.Errorf("Timeout exceeded while waiting for condition."))
		}
	}
}

func (t *Tester) WaitForTermination() error {
	timeout := time.After(t.Timeout)
	select {
	case err := <-t.doneCh:
		return t.handleError(err)
	case <-timeout:
		return t.handleError(fmt.Errorf("Timeout exceeded while waiting for termination"))
	}
}

func New(program func(tty *os.File) error) (*Tester, error) {
	doneCh := make(chan error, 1)
	defaultTimeout := time.Second

	console, err := expect.NewConsole()
	if err != nil {
		return nil, err
	}

	tester := Tester{
		doneCh:    doneCh,
		console:   console,
		lastInput: time.Now(),
		Timeout:   defaultTimeout,
		last:      "",
	}

	go func() {
		err := program(console.Tty())
		doneCh <- err
	}()

	return &tester, nil
}
