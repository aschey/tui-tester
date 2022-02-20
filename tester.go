package tester

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

type Tester struct {
	doneCh  chan struct{}
	inCh    chan []byte
	outCh   chan string
	last    string
	Timeout time.Duration
}

func (t *Tester) Read(input []byte) (n int, err error) {
	nextVal := <-t.inCh
	copied := copy(input, nextVal)

	return copied, nil
}

// from https://github.com/acarl005/stripansi/blob/master/stripansi.go
const ansi = "[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))"

var re = regexp.MustCompile(ansi)

func (t *Tester) Write(p []byte) (n int, err error) {
	strValue := string(p)
	select {
	case t.outCh <- strValue:
	default:
	}

	cleaned := strings.TrimSpace(re.ReplaceAllString(strValue, ""))
	if len(cleaned) > 0 {
		t.last = strValue
	}
	return len(p), nil
}

func (t *Tester) Send(input []byte) {
	t.inCh <- input
}

func (t *Tester) WaitFor(condition func(output string, outputLines []string) bool) (string, error) {
	timeout := time.After(t.Timeout)
	for {
		select {
		case output := <-t.outCh:
			// Send both the whole output and the output split into lines for convenience
			if condition(output, strings.Split(output, "/n")) {
				return output, nil
			}
		case <-timeout:
			return "", fmt.Errorf("Timeout exceeded")
		}

	}
}

func (t *Tester) WaitForTermination() error {
	timeout := time.After(t.Timeout)
	select {
	case <-t.doneCh:
	case <-timeout:
		return fmt.Errorf("Timeout exceeded")
	}

	return nil
}

func New(program func(tester *Tester)) Tester {
	doneCh := make(chan struct{}, 1)
	inCh := make(chan []byte, 1)
	outCh := make(chan string, 1)

	defaultTimeout := 5 * time.Second
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
