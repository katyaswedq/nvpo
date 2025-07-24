package main

import "fmt"

func isPalindrome(x int) bool {
    dano := x
    reverse := 0
    if x < 0 {
        return false
    }
    for x > 0 {
        ostatok := x % 10
        reverse = reverse*10 + ostatok
        x = x / 10
    }
    return reverse == dano
}

func main() {
    examples := []int{121, -121, 10, 12321, 0}
     for _, num := range examples {
        result := isPalindrome(num)
        fmt.Printf("Число %d является палиндромом? %t\n", num, result)
    }
}