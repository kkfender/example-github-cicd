package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	tests := EvenOrOdd(11)
	if tests != "even" {
		t.Errorf("expected even, actual %s", tests)
	}
}