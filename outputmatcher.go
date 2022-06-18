package tuitest

import (
	"strings"
	"time"

	"github.com/ActiveState/termtest/expect"
)

type outputMatcher struct {
	condition  func(state TermState) bool
	trimOutput bool
	outCh      chan TermState
	duration   *time.Duration
	now        *time.Time
}

func Matcher(condition func(state TermState) bool, trimOutput bool, outCh chan TermState, duration *time.Duration) expect.ExpectOpt {
	return func(opts *expect.ExpectOpts) error {
		opts.Matchers = append(opts.Matchers, &outputMatcher{condition: condition, trimOutput: trimOutput, outCh: outCh, duration: duration})
		return nil
	}
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
