package util

import "testing"

func AssertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("got %q, want %v", err, nil)
		t.Fail()
	}
}

func AssertEqual(t *testing.T, got, want interface{}) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
		t.Fail()
	}
}

func AssetNotNil(t *testing.T, got interface{}) {
	t.Helper()
	if got == nil {
		t.Errorf("got %q, want not nil", got)
		t.Fail()
	}
}
