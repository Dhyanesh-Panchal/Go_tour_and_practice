package main

import (
	"testing"
)

func BenchmarkReverse(b *testing.B) {
	b.ReportAllocs() // ReportAllocs() enables malloc statistics for this benchmark.
	for n := 0; n < b.N; n++ {
		Reverse("Something Something")
	}
	//fmt.Println(r)
}

func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello 123", " ", "!@#!@#1234!!Anything", "世"}

	// Provide a seed corpus
	for _, tc := range testcases {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, input string) {
		reverse := Reverse(input)
		doubleReverse := Reverse(reverse)

		if input != doubleReverse {
			t.Errorf("Input mismatch, expected %s, got %s ", input, doubleReverse)
		}
	})

}

func TestReverse(t *testing.T) {
	testcases := []struct {
		testName string
		input    string
		expected string
	}{
		{"test #1", "Hello, world", "dlrow ,olleH"},
		{"test #2", " ", " "},
		{"test #3", "!12345", "54321!"},
	}

	for _, tc := range testcases {
		t.Run(
			tc.testName, func(t *testing.T) {
				got := Reverse(tc.input)
				if got != tc.expected {
					t.Errorf("expected: %v, got: %v", tc.expected, got)
				}
			})
	}
}
