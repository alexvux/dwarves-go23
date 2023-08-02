package main

import "fmt"

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
	// create visited matrix
	rows := len(arr)
	cols := len(arr[0])
	visited := make([][]bool, rows)

	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
	}

	// count number of rectangles in arr
	count := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if !visited[i][j] {
				if arr[i][j] == 1 {
					count++
				}
				visitNeighbor(arr, visited, i, j, rows, cols)
			}
		}
	}
	fmt.Println(visited)
	return count
}

func visitNeighbor(arr [][]int, visited [][]bool, row, col, rows, cols int) {
	if row < 0 || row >= rows || col < 0 || col >= cols || visited[row][col] {
		return
	}

	visited[row][col] = true

	if arr[row][col] == 0 {
		return
	}

	visitNeighbor(arr, visited, row-1, col, rows, cols)
	visitNeighbor(arr, visited, row+1, col, rows, cols)
	visitNeighbor(arr, visited, row, col-1, rows, cols)
	visitNeighbor(arr, visited, row, col+1, rows, cols)

}
