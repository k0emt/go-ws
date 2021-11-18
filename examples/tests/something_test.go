package main

import (
	"fmt"
	"testing"
)

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
