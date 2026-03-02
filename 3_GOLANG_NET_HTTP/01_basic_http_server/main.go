package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r * http.Request){
	 if r.Method != http.MethodGet {
		http.Error(w, "Only Get is allowd", http.StatusMethodNotAllowed)
		return
	 }

	 _, err := w.Write([]byte("Hello World!"))
	 if err != nil {
		 fmt.Println("Error writing response:", err)
	 }
}

func main(){
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("try going to 8080 port ")

	err := http.ListenAndServe(":8080",nil)

	if err != nil {
		fmt.Println("error starting server: ", err)
	}

	fmt.Println(err)
}