package main

import "fmt"

func containsDuplicate(nums []int) bool {
    m := make(map[int]bool)
	for _, number := range nums{
		if m[number] {
			return true
		} 
		m[number] = true
	}
	return false
}

func main () {
	nums := []int{1, 3, 3, 4}
	fmt.Print("Найден ли дубль? - ", containsDuplicate(nums))
}