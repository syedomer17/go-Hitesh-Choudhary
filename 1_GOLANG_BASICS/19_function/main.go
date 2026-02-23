package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func SumAndProduct(a int, b int) (int , int){
	sum := a + b 
	product :=  a * b

	return sum, product
}

func main() {
	sum1 := add(10, 20)
	fmt.Println(sum1)

	sum2, product2 := SumAndProduct(5,5)
	fmt.Println(sum2, product2)

	// if you want only sum or if you want only porduct you can use _ 

	sum3 , _ := SumAndProduct(9,9)
	fmt.Println(sum3)

	_, product3 := SumAndProduct(9,9)
	fmt.Println(product3)

}