package main 

func devide(a int, b int) (div int, reminder int){
	div = a/b
	reminder =  a % b 

	return div, reminder
}

func main(){

	dev1, rem1 := devide(2,10)
	println("div:", dev1, "reminder:", rem1)

	dev2, rem2 := devide(10,3)
	println("div:", dev2, "reminder:", rem2)
}