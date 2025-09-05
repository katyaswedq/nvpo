package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Reverse interface {
	ReverseJson() string
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (p *Person) ReverseJson() string {
	reversed := map[string]interface{}{
		"age":  p.Age,
		"name": p.Name,
	}
	jsonData, err := json.Marshal(reversed)
	if err != nil {
		return "Ошибка при преобразовании JSON"
	}
	return string(jsonData)
}

func reverseHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Только POST метод разрешен", http.StatusMethodNotAllowed)
		return
	}
	person := &Person{Name: "katya", Age: 30}
	reversedJSON := person.ReverseJson()
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, reversedJSON)
}

func main() {
	person := &Person{Name: "katya", Age: 30}
	fmt.Println("Оригинальный JSON:", `{"name": "katya", "age": 30}`)
	fmt.Println("Реверсированный JSON:", person.ReverseJson())
	http.HandleFunc("/reverse", reverseHandler)
	fmt.Println("Сервер запущен на http://localhost:8080")
	fmt.Println("Отправьте POST запрос на http://localhost:8080/reverse")
	http.ListenAndServe(":8080", nil)
}