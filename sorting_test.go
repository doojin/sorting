package sorting

import (
	"testing"
	"reflect"
)

type TestCase struct {
	input []int
	expected []int
}

var testCases []TestCase = []TestCase {
	// Positive numbers
	TestCase {
		input: []int {9, 8, 7, 6, 5, 4, 3, 2, 1},
		expected: []int {1, 2, 3, 4, 5, 6, 7, 8, 9},
	},
	
	// Negative numbers
	TestCase {
		input: []int {-6, -3, -2, -9, -7, -3},
		expected: []int {-9, -7, -6, -3, -3, -2},
	},
	
	// Positive and negative numbers
	TestCase {
		input: []int {5, -1, 8, 0, -2, 3},
		expected: []int {-2, -1, 0, 3, 5, 8},
	},
	
	// Decreasing sorted slice
	TestCase {
		input: []int {5, 4, 3, 2, 1, 0, -1, -2},
		expected: []int {-2, -1, 0, 1, 2, 3, 4, 5},
	},
	
	// Empty slice
	TestCase {
		input: []int {},
		expected: []int {},
	},
}

func TestBubbleSort(t *testing.T) {
	for _, testCase := range(testCases) {
		output := BubbleSort(testCase.input)
		if	!reflect.DeepEqual(testCase.expected, output) {
			t.Error("Bubble sort test error. Expected:", testCase.expected, "got:", output)
		}
	}
}