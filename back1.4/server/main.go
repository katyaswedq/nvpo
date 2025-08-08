package main

import (
	"fmt"
	"net/http"
	"time"
)

func main () {
	http.HandleFunc("/time", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, time.Now().Format("Время: 2006-01-02 15:04:05"))
	})
	fmt.Print("Сервер запущен: http://localhost:8080/time")
	http.ListenAndServe(":8080", nil)
}