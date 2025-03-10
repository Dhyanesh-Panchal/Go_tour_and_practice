package fib_test

import (
	"benchmarking/fib"
	"fmt"
	"testing"
)

func BenchmarkFibonacci(b *testing.B) {
	inputs := []uint64{1, 5, 10, 15, 20}

	for _, input := range inputs {
		b.Run(fmt.Sprintf("benchmark-input-%d", input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fib.Fibonacci(input)
			}
		})
	}
}

func BenchmarkQuickFibonacci(b *testing.B) {
	inputs := []uint64{1, 5, 10, 15, 20}

	for _, input := range inputs {
		b.Run(fmt.Sprintf("benchmark-input-%d", input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fib.QuickFibonacci(input)
			}
		})
	}
}

func TestFibonacci(t *testing.T) {
	testcases := []struct {
		input, expected uint64
	}{
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
	}

	for _, tc := range testcases {
		t.Run("test fibonacci", func(t *testing.T) {
			//t.Parallel() // run in parallel
			got := fib.Fibonacci(tc.input)
			if got != tc.expected {
				t.Errorf("Fibonacci(%d): expected %d, got %d", tc.input, tc.expected, got)
			}
		})

	}
}

func TestQuickFibonacci(t *testing.T) {
	testcases := []struct {
		input, expected uint64
	}{
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
	}

	for _, tc := range testcases {
		t.Run("test fibonacci", func(t *testing.T) {
			//t.Parallel() // run in parallel
			got := fib.QuickFibonacci(tc.input)
			if got != tc.expected {
				t.Errorf("Fibonacci(%d): expected %d, got %d", tc.input, tc.expected, got)
			}
		})

	}
}
