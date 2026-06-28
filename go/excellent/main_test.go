package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	tests := EvenOrOdd(10)
	if tests != "even" {
		t.Errorf("expected even, actual %s", tests)
	}
}