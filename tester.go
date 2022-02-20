package tester

import (
	"regexp"
	"strings"
)

type Tester struct {
	doneCh chan struct{}
	inCh   chan []byte
	outCh  chan string
	last   string
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

func (t *Tester) WaitFor(condition func(output string) bool) string {
	for output := range t.outCh {
		if condition(output) {
			return output
		}
	}
	return ""
}

func (t *Tester) WaitForTermination() {
	<-t.doneCh
}

func New(program func(tester *Tester)) Tester {
	doneCh := make(chan struct{}, 1)
	inCh := make(chan []byte, 1)
	outCh := make(chan string, 1)
	tester := Tester{doneCh: doneCh, inCh: inCh, outCh: outCh, last: ""}

	go func() {
		program(&tester)
		doneCh <- struct{}{}
	}()

	return tester
}