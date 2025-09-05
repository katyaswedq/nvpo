package main

import "fmt"

type User struct {
	Name string
	Age  int
	Next *User
}

func (u *User) Add(name string, age int) {
	newUser := &User{Name: name, Age: age}
	if u.Name == "" && u.Age == 0 && u.Next == nil {
		u.Name = name
		u.Age = age
		return
	}
	copy := u
	for copy.Next != nil {
		copy = copy.Next
	}
	copy.Next = newUser
}

func (u *User) Remove(name string) {
	if u.Name == name {
		if u.Next != nil {
			u.Name = u.Next.Name
			u.Age = u.Next.Age
			u.Next = u.Next.Next
		} else {
			u.Name = ""
			u.Age = 0
			u.Next = nil
		}
		return
	}
	copy := u
	for copy.Next != nil {
		if copy.Next.Name == name {
			copy.Next = copy.Next.Next
			return
		}
		copy = copy.Next
	}
}

func main() {
	user := &User{}
	user.Add("Alice", 25)
	fmt.Println(user.Name, user.Age, user.Next)
	user.Add("Bob", 30)
	fmt.Println(user.Name, user.Age, user.Next)
	user.Remove("Bob")
	fmt.Println(user.Name, user.Age, user.Next)
}
