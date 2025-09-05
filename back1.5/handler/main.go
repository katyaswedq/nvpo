package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello from custom handler!")
}

func main() {
	fmt.Print("Сервер запущен: http://localhost:8080")
    http.ListenAndServe(":8080", MyHandler{})
}