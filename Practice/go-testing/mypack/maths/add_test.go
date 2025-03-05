package maths_test

import (
	"mypack/maths"
	"testing"
)

type adderTestCase struct {
	testName       string
	a, b, expected int
}

var tests = []adderTestCase{
	{"Test Case: #1", 1, 2, 3},
	{"Test Case: #2", 3, 2, 5},
	{"Test Case: #3", 0, 4, 4},
	{"Test Case: #4", -5, 4, -1},
}

//func TestAdder(t *testing.T) {
//	for _, tc := range tests {
//		got := maths.Adder(tc.a, tc.b)
//
//		if got != tc.expected {
//			t.Errorf("Expected %d, got %d", tc.expected, got)
//		}
//	}
//}

// Using t.Run()
func TestAdder(t *testing.T) {
	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {
			got := maths.Adder(tc.a, tc.b)

			if got != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, got)
			}
		})

	}
}
