package tests_test

import (
	"test_examples/race"
	"testing"
)

func TestGetCounter(t *testing.T) {
	counter := race.GetCounter()
	if counter != 50000 {
		t.Error("unexpected counter", counter)
	}
}
