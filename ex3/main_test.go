package main

import (
	"testing"
)

func Test_countRectangles(t *testing.T) {
	testCases := []struct {
		name     string
		input    [][]int
		expected int
	}{
		{
			name: "valid input 1",
			input: [][]int{ // 8 rows 7 cols
				{1, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{1, 0, 0, 1, 1, 1, 0},
				{0, 1, 0, 1, 1, 1, 0},
				{0, 1, 0, 0, 0, 0, 0},
				{0, 1, 0, 1, 1, 0, 0},
				{0, 0, 0, 1, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 1},
			},
			expected: 6,
		},
		{
			name: "valid input 2",
			input: [][]int{ // 8 rows 7 cols
				{1, 1, 0, 0, 0, 0, 1},
				{1, 1, 0, 0, 0, 0, 0},
				{0, 0, 0, 1, 1, 0, 0},
				{0, 1, 0, 1, 1, 0, 0},
				{0, 1, 0, 0, 0, 1, 1},
				{0, 1, 0, 1, 1, 0, 0},
				{1, 0, 0, 1, 1, 0, 0},
				{1, 0, 0, 0, 0, 0, 1},
			},
			expected: 8,
		},
		{
			name: "adjacent rectangle left",
			input: [][]int{ // 8 rows 7 cols
				{1, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 1, 0, 0},
				{1, 0, 0, 1, 1, 1, 0},
				{0, 1, 0, 1, 1, 1, 0},
				{0, 1, 0, 0, 0, 0, 0},
				{0, 1, 0, 1, 1, 0, 0},
				{0, 0, 0, 1, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 1},
			},
			expected: -1,
		},
		{
			name: "adjacent rectangle right",
			input: [][]int{ // 8 rows 7 cols
				{1, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 1, 1, 0, 0},
				{1, 0, 0, 1, 1, 1, 0},
				{0, 1, 0, 1, 1, 1, 0},
				{0, 1, 0, 0, 0, 0, 0},
				{0, 1, 0, 1, 1, 0, 0},
				{0, 0, 0, 1, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 1},
			},
			expected: -1,
		},
		{
			name: "invalid rectangle border",
			input: [][]int{ // 8 rows 7 cols
				{1, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 1, 1, 1, 0},
				{1, 0, 0, 1, 1, 1, 0},
				{0, 1, 0, 1, 1, 1, 0},
				{0, 1, 0, 1, 1, 1, 0},
				{0, 1, 0, 1, 1, 0, 0},
				{0, 0, 0, 1, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 1},
			},
			expected: -1,
		},
		{
			name: "invalid rectangle inside",
			input: [][]int{ // 8 rows 7 cols
				{1, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 1, 1, 1, 0},
				{1, 0, 0, 1, 0, 1, 0},
				{0, 1, 0, 1, 1, 1, 0},
				{0, 1, 0, 0, 0, 0, 0},
				{0, 1, 0, 1, 1, 0, 0},
				{0, 0, 0, 1, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 1},
			},
			expected: -1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := countRectangles(tc.input)
			if actual != tc.expected {
				t.Errorf("Test %s: expected %d, got %d", tc.name, tc.expected, actual)
			}
		})
	}
}
