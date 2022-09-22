package race_test

import (
  "testing"
  "test_examples/race"
)

func TestGetCounter(t *testing.T) {
  counter := race.GetCounter()
  if counter != 50000 {
    t.Error("unexpected counter", counter)
  }
}