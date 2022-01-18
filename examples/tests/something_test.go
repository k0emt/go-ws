package main

import (
	"fmt"
	"testing"
)

func TestDoSomethingNoError(t *testing.T) {
	result, _ := doSomething("AAAbbbCCC")
	expectedValue := "AAABBBCCC"

	if result != expectedValue {
		t.Errorf("doSomething(AAAbbbCCC) = %s, expected %s", result, expectedValue)
	}
}

func TestDoSomethingError(t *testing.T) {
	_, error := doSomething("")
	// expectedValue := "Should have provided a non-empty string"
	expected := "need to provide text to operate on"

	if error == nil {
		t.Errorf("no error")
	} else if *error != expected {
		t.Errorf("Expected %s, Received %s", expected, *error)
	}
}

func TestDoSomething(t *testing.T) {
	val, err := doSomething("aaaaAAAA")

	if err != nil {
		t.Fatalf("shouldn't have an error %v", err)
	}

	if val != "AAAAAAAA" {
		t.Fatalf("Expected: AAAAAAAA - received: %v", val)
	}
}

type testCase struct {
	input    string
	expected string
}

func TestMultipleDoSomething(t *testing.T) {
	testCases := []testCase{
		{"eee", "EEE"},
		{"I", "I"},
		{"oH", "OH"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s ", tc.input), func(t *testing.T) {
			converted, err := doSomething(tc.input)

			if err != nil {
				t.Fatalf("shouldn't have an error %v", err)
			}

			if tc.expected != converted {
				t.Fatalf("Expected: %s - received: %v", tc.expected, converted)
			}
		})
	}
}

func BenchmarkDoSomething(b *testing.B) {
	someText := []string{"alpha", "bRavo", "Charlie", "DELTA", "echO", "FoxTrot", "GOLF",
		"HOTel"}

	for i := 0; i < b.N; i++ {
		doSomething(someText[i%len(someText)])
	}
}
