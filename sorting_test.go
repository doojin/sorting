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
	TestCase {
		input: []int {5, -1, 8, 5, 3},
		expected: []int {-1, 3, 5, 5, 8},
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