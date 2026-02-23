package main

import "fmt"

type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

func main() {
	u1 := User{
		ID:    1,
		Name:  "Syed Omer Ali",
		Email: "syedomerali2006@gmail.com",
		Age:   19,
	}

	fmt.Println(u1,u1.ID, u1.Email,u1.Name)

	u1.Age = 20 
	fmt.Println(u1.Age)

	u2 := User{
	ID:    2,
	Name:  "John Doe",
	}

	fmt.Println(u2)
}