package testutil

import (
	"testing"
)

func AssertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatalf("%s != %s", a, b)
	}
}

func AssertNil(t *testing.T, a interface{}) {
	if a != nil {
		t.Fatalf("%s should be nil", a)
	}
}

func AssertNotNil(t *testing.T, a interface{}) {
	if a == nil {
		t.Fatalf("%s should not be nil", a)
	}
}

func AssertTrue(t *testing.T, a interface{}) {
	if a != true {
		t.Fatalf("%s should not be false", a)
	}
}
