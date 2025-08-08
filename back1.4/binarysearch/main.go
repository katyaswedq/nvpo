package main 

import "fmt"

func binarySearch(arr []int, target int) int {
	start := 0
	finish := len(arr) - 1

	for start <= finish {
		middle := start + (finish - start) / 2
		if arr[middle] == target {
			return middle
		} else if arr[middle] < target {
			start = middle + 1
		} else {
			finish = middle - 1
		}
	}
	return -1
}

func main () {
	array := []int{1, 3, 5, 8, 13, 45}
	target := 3
	index := binarySearch(array, target)
	if index != -1 {
		fmt.Printf("Элемент %v найден под индексом %v", target, index)
	} else {
		fmt.Print("Элемент не найден")
	}
}