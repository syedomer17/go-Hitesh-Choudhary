package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CatFactResponse struct {
	Fact string `json:"fact"`
	Length int `json:"length"`
}

func main(){
	url := "https://catfact.ninja/facts"

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("error", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		fmt.Println(resp.Status)
		return
	}

	bodyBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("error", err)
		return
	}

	var data CatFactResponse 

	if err := json.Unmarshal(bodyBytes,&data); err != nil {
		fmt.Println("error", err)
		return
	}

	fmt.Println(data.Fact,data.Length,len(bodyBytes))

}