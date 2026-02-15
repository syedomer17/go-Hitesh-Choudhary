package main 

import (
	"fmt"
)

func main(){
	isLoggedIn := true 
	isAdmin := false 
	hasSubscription := true 

	canOpenDashboard := isLoggedIn && hasSubscription

	canDeletePost := isAdmin || (isLoggedIn && hasSubscription)

	fmt.Println(canOpenDashboard, canDeletePost)
}