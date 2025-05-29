package base

import (
	"testing"
)

type TestCase struct {
	Name     string
	Input    []float64
	Expected float64
}

var s SimpleMathService

func runTest(t *testing.T, test TestCase, f func([]float64) float64) {
	t.Run(test.Name, func(t *testing.T) {
		result := f(test.Input)
		if result != test.Expected {
			t.Errorf("Expected %v, got %v", test.Expected, result)
		}
	})
}

func TestBaseAdd(t *testing.T) {
	tests := []TestCase{
		{
			Name:     "[] should be 0",
			Input:    []float64{},
			Expected: 0,
		},
		{
			Name:     "[x] should be x",
			Input:    []float64{1},
			Expected: 1,
		},
		{
			Name:     "[x, y] should be x+y",
			Input:    []float64{1, 2},
			Expected: 3,
		},
	}

	for _, test := range tests {
		runTest(t, test, s.Add)
	}
}

func TestBaseSub(t *testing.T) {
	tests := []TestCase{
		{
			Name:     "[] should be 0",
			Input:    []float64{},
			Expected: 0,
		},
		{
			Name:     "[x] should be x",
			Input:    []float64{1},
			Expected: 1,
		},
		{
			Name:     "[x, y] should be x-y",
			Input:    []float64{1, 2},
			Expected: -1,
		},
	}

	for _, test := range tests {
		runTest(t, test, s.Sub)
	}
}

func TestBaseMul(t *testing.T) {
	tests := []TestCase{
		{
			Name:     "[] should be 0",
			Input:    []float64{},
			Expected: 0,
		},
		{
			Name:     "[x] should be x",
			Input:    []float64{1},
			Expected: 1,
		},
		{
			Name:     "[x, y] should be x*y",
			Input:    []float64{4, 2},
			Expected: 8,
		},
	}

	for _, test := range tests {
		runTest(t, test, s.Mul)
	}
}

func TestBaseDiv(t *testing.T) {
	tests := []TestCase{
		{
			Name:     "[] should be 0",
			Input:    []float64{},
			Expected: 0,
		},
		{
			Name:     "[x] should be x",
			Input:    []float64{1},
			Expected: 1,
		},
		{
			Name:     "[x, y] should be x/y",
			Input:    []float64{4, 2},
			Expected: 2,
		},
	}

	for _, test := range tests {
		runTest(t, test, s.Div)
	}
}
