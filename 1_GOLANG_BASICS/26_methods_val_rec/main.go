package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	u1 := User{
		Name: "John",
		Age:  30,
	}
	fmt.Println(u1.Intro())
}

func (u User) Intro() string {
	return fmt.Sprintf("My name is %s and I am %d years old.", u.Name, u.Age)
}