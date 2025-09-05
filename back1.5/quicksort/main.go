package main

import "fmt"

func quickSort(arr []int) []int {
    if len(arr) < 2 {
        return arr
    }
    pivot := arr[len(arr)/2]
    var less, equal, greater []int
    for _, num := range arr {
        switch {
        case num < pivot:
            less = append(less, num)
        case num == pivot:
            equal = append(equal, num)
        default:
            greater = append(greater, num)
        }
    }
    left := quickSort(less)
    right := quickSort(greater)
    result := append(left, equal...)
    result = append(result, right...)
    return result
}

func main () {
	arr := []int{45, 6, 1, 90, 33}
	fmt.Print(quickSort(arr))
}