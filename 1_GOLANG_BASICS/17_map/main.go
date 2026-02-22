package main

import "fmt"

func main() {
	// map[keyType]valueType
	ages := map[string]int{
		"Omer": 19,
		"Ali":  20,
	}

	fmt.Println(ages["Omer"], len(ages))

	var scores map[string]int // this is an nil map 

	fmt.Println(scores)

	scores = make(map[string]int)

	scores["Math"] = 90 

	fmt.Println(scores)

	users := map[string]string{
		"user1": "Omer",
		"user2": "Ali",
	}
	fmt.Println(users)

	// to delete an key value pair from a map we can use the delete function

	delete(users, "user2")
	fmt.Println(users)
}