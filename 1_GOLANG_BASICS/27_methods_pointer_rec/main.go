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
	fmt.Println(u1.Name)
	fmt.Println("Before:u1.Age =", u1.Age)
	u1.Birthday()
	fmt.Println("After:u1.Age =", u1.Age)
}

func ( u *User) Birthday() {
	u.Age++
}