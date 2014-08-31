package sorting

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	input    []int
	expected []int
}

var testCases = []TestCase{
	// Positive numbers
	TestCase{
		input:    []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
		expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
	},

	// Negative numbers
	TestCase{
		input:    []int{-6, -3, -2, -9, -7, -3},
		expected: []int{-9, -7, -6, -3, -3, -2},
	},

	// Positive and negative numbers
	TestCase{
		input:    []int{5, -1, 8, 0, -2, 3},
		expected: []int{-2, -1, 0, 3, 5, 8},
	},

	// Decreasing sorted slice
	TestCase{
		input:    []int{5, 4, 3, 2, 1, 0, -1, -2},
		expected: []int{-2, -1, 0, 1, 2, 3, 4, 5},
	},

	// Empty slice
	TestCase{
		input:    []int{},
		expected: []int{},
	},
}

func TestStupudSort(t *testing.T) {
	for _, testCase := range testCases {
		output := StupidSort(testCase.input)
		assert.Equal(t, testCase.expected, output)
	}
}

func TestBubbleSort(t *testing.T) {
	for _, testCase := range testCases {
		output := BubbleSort(testCase.input)
		assert.Equal(t, testCase.expected, output)
	}
}

func TestSelectionSort(t *testing.T) {
	for _, testCase := range testCases {
		output := SelectionSort(testCase.input)
		assert.Equal(t, testCase.expected, output)
	}
}

func TestCoctailSort(t *testing.T) {
	for _, testCase := range testCases {
		output := CoctailSort(testCase.input)
		assert.Equal(t, testCase.expected, output)
	}
}

func TestGnomeSort(t *testing.T) {
	for _, testCase := range testCases {
		output := GnomeSort(testCase.input)
		assert.Equal(t, testCase.expected, output)
	}
}

func TestInsertionSort(t *testing.T) {
	for _, testCase := range testCases {
		output := InsertionSort(testCase.input)
		assert.Equal(t, testCase.expected, output)
	}
}

func TestQuickSort(t *testing.T) {
	for _, testCase := range testCases {
		output := QuickSort(testCase.input, 0, len(testCase.input)-1)
		assert.Equal(t, testCase.expected, output)
	}
}

func TestMergeSort(t *testing.T) {
	for _, testCase := range testCases {
		output := MergeSort(testCase.input)
		assert.Equal(t, testCase.expected, output)
	}
}

func TestTreeSort(t *testing.T) {
	for _, testCase := range testCases {
		output := TreeSort(testCase.input)
		assert.Equal(t, testCase.expected, output)
	}
}
