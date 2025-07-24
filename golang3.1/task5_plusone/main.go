package main

import "fmt"

func plusOne(digits []int) []int {
    n := len(digits)
    for i := n - 1; i >= 0; i-- {
        if digits[i] != 9 {
            digits[i]++
            return digits
        } else {
            digits[i] = 0
            if i == 0 && digits[i] == 0 {
                digits = append([]int{1}, digits...)
                return digits
            }
        }
    }
    return digits
}

func main() {
    testCases := [][]int{
        {1, 2, 3},
        {4, 3, 2, 1},
        {9},
        {9, 9},
        {0},
    }

    for _, tc := range testCases {
        fmt.Println(plusOne(tc))
    }
}