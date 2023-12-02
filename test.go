package main

import "testing"

func AssertEqual[T comparable](t *testing.T, expect, result T) {
	t.Helper()

	if result != expect {
		t.Errorf("expected %v, got %v", expect, result)
	}
}
