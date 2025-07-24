package main

import (
	"fmt"
	"strings"
)

func isValid(s string) bool {
	for {
		oldLen := len(s)
		s = strings.ReplaceAll(s, "()", "")
		s = strings.ReplaceAll(s, "[]", "")
		s = strings.ReplaceAll(s, "{}", "")
		if len(s) == oldLen {
			break
		}
	}
	return len(s) == 0
}

func main() {
	examples := []string{
		"()",        
		"()[]{}",    
		"(]",        
		"([)]",      
		"{[]}",      
		"[[[[",      
		"",  
	}
	for _, example := range examples {
		fmt.Println(isValid(example))
	}
}