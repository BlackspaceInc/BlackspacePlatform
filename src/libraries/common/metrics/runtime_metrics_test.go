package metrics_test

import (
	"testing"

	"github.com/BlackspaceInc/common/metrics"
)

func TestRuntimeCountersPresent(t *testing.T) {
	t.Run("TestName:NumberOfGoRoutinesCounterExists", TestNumberOfGoRoutinesCounterExists)
}

func TestNumberOfGoRoutinesCounterExists(t *testing.T) {
	var counterName = metrics.NumberOfGoRoutines
	validateGaugeFuncExists(t, counterName)
}
