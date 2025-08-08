package main

import (
	"fmt"
	"log"
	"os"
)

func main () {
	poema, err := os.ReadFile("file.txt")
	if err != nil {
		log.Fatal("Ошибка:", err)
	}
	fmt.Print(string(poema))
}