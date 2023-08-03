package main

import (
	"fmt"
)

func main() {
	arr := [][]int{ // 8 rows 7 cols
		{1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 1, 1, 1, 0},
		{0, 1, 0, 1, 1, 1, 0},
		{0, 1, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 1, 0, 0},
		{0, 0, 0, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 1},
	}

	count := countRectangles(arr)
	fmt.Println(count)
}

func countRectangles(arr [][]int) int {
	count := 0
	// init visited matrix
	rows := len(arr)
	cols := len(arr[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			// found a rectangle
			if arr[row][col] == 1 && !visited[row][col] {
				// get size of rectangle
				nStart, nEnd := col, col
				for nEnd < cols && arr[row][nEnd] == 1 {
					nEnd++
				}
				mStart, mEnd := row, row
				for mEnd < rows && arr[mEnd][col] == 1 {
					mEnd++
				}
				// check if rectangle has cell with value '0' or has been visited
				if !isValidRectangle(arr, visited, mStart, mEnd, nStart, nEnd) {
					return -1
				}
				// check if rectangle has no adjacent rectangles or not
				if !hasNoAdjacentRectangle(arr, mStart, mEnd, nStart, nEnd, rows, cols) {
					return -1
				}
				count++
			}
		}
	}
	return count
}

func isValidRectangle(arr [][]int, visited [][]bool, rowStart, rowEnd, colStart, colEnd int) bool {
	for row := rowStart; row < rowEnd; row++ {
		for col := colStart; col < colEnd; col++ {
			if arr[row][col] == 0 || visited[row][col] {
				return false
			}
			visited[row][col] = true
		}
	}
	return true
}

func hasNoAdjacentRectangle(arr [][]int, rowStart, rowEnd, colStart, colEnd, rows, cols int) bool {
	// check top border
	if rowStart > 0 {
		for col := colStart; col < colEnd; col++ {
			if arr[rowStart-1][col] != 0 {
				return false
			}
		}
	}
	// check bottom border
	if rowEnd < rows-1 {
		for col := colStart; col < colEnd; col++ {
			if arr[rowEnd][col] != 0 {
				return false
			}
		}
	}
	// check left border
	if colStart > 0 {
		for row := rowStart; row < rowEnd; row++ {
			if arr[row][colStart-1] != 0 {
				return false
			}
		}
	}
	// check right border
	if colEnd < cols-1 {
		for row := rowStart; row < rowEnd; row++ {
			if arr[row][colEnd] != 0 {
				return false
			}
		}
	}
	return true
}
