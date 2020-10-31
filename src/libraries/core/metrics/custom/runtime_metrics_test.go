package custom_test

import (
	"testing"
)

func TestRuntimeCountersPresent(t *testing.T) {
	t.Run("TestName:NumberOfGoRoutinesCounterExists", TestNumberOfGoRoutinesCounterExists)
}

func TestNumberOfGoRoutinesCounterExists(t *testing.T) {
	var counterName = NumberOfGoRoutines
	validateGaugeFuncExists(t, counterName)
}
