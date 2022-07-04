package tuitest

import (
	"fmt"
	"time"
)

type Option func(t *Tester) error

func WithDefaultWaitTimeout(timeout time.Duration) Option {
	return func(t *Tester) error {
		t.defaultWaitTimeout = timeout
		return nil
	}
}

func WithMinInputInterval(minInterval time.Duration) Option {
	return func(t *Tester) error {
		t.minInputInterval = minInterval
		return nil
	}
}

func WithTerminationTimeout(terminationTimeout time.Duration) Option {
	return func(t *Tester) error {
		t.terminationTimeout = terminationTimeout
		return nil
	}
}

func WithErrorHandler(onError func(err error) error) Option {
	return func(t *Tester) error {
		if onError == nil {
			return fmt.Errorf("onError can't be nil")
		}
		t.onError = onError
		return nil
	}
}
