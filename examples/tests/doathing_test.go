package main

import "testing"

func TestAthing(t *testing.T) {
	actual, ok := doAthing("SOMETHING")

	if actual != "something" || !ok {
		t.Errorf("Expected something, true, got: %s %t", actual, ok)
	}

}

func TestAthingSomeMore(t *testing.T) {
	actual, ok := doAthing("")

	if actual != "" || ok {
		t.Errorf("Expected \"\", false, got: %s %t", actual, ok)
	}

}
