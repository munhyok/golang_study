/*
	Slice of Slices

	2차원 배열...

	rows := [][]int{}


*/

package main

import "fmt"

func createMatrix(rows, cols int) [][]int {
	// ?
	matrix := [][]int{}

	for i := 0; i < rows; i++ {
		row := []int{}
		for j := 0; j < cols; j++ {
			row = append(row, i*j)
		}
		matrix = append(matrix, row)

	}
	fmt.Println(matrix) // [[0 0 0] [0 1 2] [0 2 4]]
	return matrix

}

// dont edit below this line

func test(rows, cols int) {
	fmt.Printf("Creating %v x %v matrix...\n", rows, cols)
	matrix := createMatrix(rows, cols)
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
	fmt.Println("===== END REPORT =====")
}

func main() {
	test(3, 3)
	//test(5, 5)
	//test(10, 10)
	//test(15, 15)
}
