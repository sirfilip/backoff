package backoff

import (
	"errors"
	"testing"
	"time"
)

func TestFailure(t *testing.T) {
	result := 0
	err := Retry(func() error {
		result = result + 1
		if result < 6 {
			return errors.New("Failure")
		}
		return nil
	}, 5, time.Millisecond)
	if err == nil {
		t.Errorf("Expected an error but got nil insted")
	}
}

func TestScuccess(t *testing.T) {
	result := 0
	err := Retry(func() error {
		result = result + 1
		if result < 6 {
			return errors.New("Failure")
		}
		return nil
	}, 7, time.Millisecond)
	if err != nil {
		t.Errorf("Expected nil but got: %v", err)
	}
}
