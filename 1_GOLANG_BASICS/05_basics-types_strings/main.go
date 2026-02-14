package main 

import ( 
	"fmt"
	"strings"
)

func main(){
	firstName := "Syed Omer"
	lastName := "Ali"

	fullName := firstName + " " + lastName

	fmt.Println(fullName)

	fmt.Println(strings.ToUpper(fullName))
}