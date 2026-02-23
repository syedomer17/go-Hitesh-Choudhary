package main 

func sumAll(numbs ...int) int {
	total := 0

	for _, currentValue := range numbs {
		total = total + currentValue
	}
	return total
}

func main(){
	println(sumAll(1,2,3,4,5))

	values := []int{10,20}
	println((sumAll(values...)))

	res := func(n int) int{
		return n * 2
	}
	println(res(5))

	res2 := func(a int, b int) int {
	return a * b
	} 

	println(res2(5,10))
}