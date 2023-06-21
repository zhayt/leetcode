package main

import "fmt"

func main() {
	var s = User{"test"}
	s.PrintName()
	s.UpdateName()
	s.PrintName()
}

type User struct {
	Name string
}

func (m User) UpdateName() {
	m.Name = "new name"
}

func (m User) PrintName() {
	fmt.Println("my name =", m.Name)
}
