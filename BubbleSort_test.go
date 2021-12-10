package main

import (
	"reflect"
	"testing"
)

func TestArr(t *testing.T) {
	// Arrange
	testTable := []struct {
		array    []int
		expected []int
	}{
		{
			array:    []int{23, 28, 88, 99, 76, 67, 9, 79, 92, 20, 33, 94, 74, 24, 66, 54, 37, 0, 51, 19},
			expected: []int{0, 9, 19, 20, 23, 24, 28, 33, 37, 51, 54, 66, 67, 74, 76, 79, 88, 92, 94, 99},
		},
		{
			array:    []int{9, 7, 3, 8, 1, 0, 0, 100},
			expected: []int{0, 0, 1, 3, 7, 8, 9, 100},
		},
		{
			array:    []int{},
			expected: []int{},
		},
	}

	for _, testCase := range testTable {
		//Act
		BubbleSort(testCase.array)
		//Assert
		if reflect.DeepEqual(testCase.expected, testCase.array) == false {
			t.Errorf("\n Incorrect result.\n Expected %d, \n Result: %d , \n", testCase.expected, testCase.array)
		}
	}

}
