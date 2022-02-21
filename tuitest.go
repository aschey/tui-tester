package tuitest

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

type Tester struct {
	doneCh     chan struct{}
	inCh       chan []byte
	outCh      chan string
	last       string
	Timeout    time.Duration
	TrimOutput bool
	RemoveAnsi bool
}

func (t *Tester) Read(input []byte) (n int, err error) {
	nextVal := <-t.inCh
	copied := copy(input, nextVal)

	return copied, nil
}

// from https://github.com/acarl005/stripansi/blob/master/stripansi.go
const ansi = "[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))"

var re = regexp.MustCompile(ansi)

func RemoveAnsiSequences(str string) string {
	return re.ReplaceAllString(str, "")
}

func (t *Tester) Write(p []byte) (n int, err error) {
	strValue := string(p)
	select {
	case t.outCh <- strValue:
	default:
	}

	cleaned := strings.TrimSpace(RemoveAnsiSequences(strValue))
	if len(cleaned) > 0 {
		t.last = strValue
	}
	return len(p), nil
}

func (t *Tester) SendBytes(input []byte) {
	t.inCh <- input
}

func (t *Tester) SendByte(input byte) {
	t.SendBytes([]byte{input})
}

func (t *Tester) SendString(input string) {
	t.SendBytes([]byte(input))
}

func (t *Tester) WaitFor(condition func(output string, outputLines []string) bool) (string, []string, error) {
	timeout := time.After(t.Timeout)
	last := ""
	for {
		select {
		case output := <-t.outCh:
			if t.RemoveAnsi {
				output = RemoveAnsiSequences(output)
			}
			if t.TrimOutput {
				output = strings.TrimSpace(output)
			}
			// Send both the whole output and the output split into lines for convenience
			outputLines := strings.Split(output, "\n")
			last = strings.TrimSpace(RemoveAnsiSequences(output))
			if condition(output, outputLines) {
				return output, outputLines, nil
			}
		case <-timeout:
			return "", []string{}, fmt.Errorf("Timeout exceeded while waiting for condition. Last returned output: %s", last)
		}

	}
}

func (t *Tester) WaitForTermination() error {
	timeout := time.After(t.Timeout)
	select {
	case <-t.doneCh:
	case <-timeout:
		return fmt.Errorf("Timeout exceeded while waiting for termination")
	}

	return nil
}

func New(program func(tester *Tester)) Tester {
	doneCh := make(chan struct{}, 1)
	inCh := make(chan []byte, 1)
	outCh := make(chan string, 1)

	defaultTimeout := 1 * time.Second
	tester := Tester{
		doneCh:  doneCh,
		inCh:    inCh,
		outCh:   outCh,
		Timeout: defaultTimeout,
		last:    "",
	}

	go func() {
		program(&tester)
		doneCh <- struct{}{}
	}()

	return tester
}
