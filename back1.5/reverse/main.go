package main

import "fmt"

func reverse(x int) int {
    res := 0
    for x != 0 {
        ost := x % 10
        res = res * 10 + ost
        x = x / 10 
    }
    return res
}

func main () {
	number := 632
	fmt.Print(reverse(number))
}