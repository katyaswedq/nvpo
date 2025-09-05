package main

import "fmt"

func generate(numRows int) [][]int {
    matrix := make([][]int, numRows)
    for i := range matrix {
        matrix[i] = make([]int, i+1)
        matrix[i][0], matrix[i][i] = 1, 1
        for j := 1; j < i; j++ {
            matrix[i][j] = matrix[i-1][j-1] + matrix[i-1][j]
        }
    }
    return matrix
}

func main () {
	num := 5
	fmt.Print(generate(num))
}