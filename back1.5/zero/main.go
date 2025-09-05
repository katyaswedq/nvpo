package main

import "fmt"

func moveZeroes(nums []int) {
    zero := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != 0 {
            nums[zero], nums[i] = nums[i], nums[zero]
            zero++
        }
    }
}

func main () {
	nums := []int{0, 1, 5, 0, 6, 3}
	moveZeroes(nums)
	fmt.Print(nums)
}