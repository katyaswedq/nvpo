package main

import "fmt"

func romanToInt(s string) int {
    m := map[string]int{
        "I": 1,
        "V": 5,
        "X": 10,
        "L": 50,
        "C": 100,
        "D": 500,
        "M": 1000,
    }
    total := 0
    for i := 0; i < len(s); i++ {
        x := m[string(s[i])]
        if i+1 < len(s) && m[string(s[i+1])] > x {
            total = total - x
        } else {
            total = total + x
        }
    }
    return total
}

func main() {
    examples := []string{"III", "IV", "IX", "LVIII", "MCMXCIV"}
    for _, roman := range examples {
        result := romanToInt(roman)
        fmt.Printf("Римское число %s = %d\n", roman, result)
    }
}