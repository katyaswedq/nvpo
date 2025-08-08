package main

import "fmt"

func bubbleSort(arr []int) {
    x := len(arr)
    for i := 0; i < x-1; i++ {
        for j := 0; j < x-i-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
}

func main() {
    data := []int{13, 8, 56, 43, 75}
    fmt.Println("До:", data)
    bubbleSort(data)
    fmt.Println("После:", data)
}