package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type CatFactResponse struct {
	Fact string `json:"fact"`
	Length int `json:"length"`
}

func writeJson(w http.ResponseWriter, statusCode int, data any){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	_ = json.NewEncoder(w).Encode(data)
}

func fetcCatFact() (*CatFactResponse, error){
	url := "https://catfact.ninja/fact"

	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		return nil, fmt.Errorf("external api failed : %s", resp.Status)
	}

	bodyBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var data CatFactResponse 

	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func externalHandler(w http.ResponseWriter, r *http.Request){

	if r.Method != http.MethodGet {
		writeJson(w, http.StatusMethodNotAllowed, map[string]any{
			"ok":false,
			"error":"only get method allwoed",
		})
		return
	}

	data, err := fetcCatFact()

	if err != nil {
		writeJson(w, http.StatusBadGateway, map[string]any{
			"ok":false,
			"error":"failed to fetch",
		})
		return
	}

	writeJson(w, http.StatusOK, map[string]any{
		"ok": true,
		"timeStamp": time.Now().UTC(),
		"external": map[string]any{
			"source":"CatFact.ninja",
			"fact": data.Fact,
			"length":data.Length,
		},
	})
}

func main() {

	http.HandleFunc("/external", externalHandler)
	err := http.ListenAndServe(":5000",nil)

	if err != nil {
		fmt.Println("error", err)
		return
	}
}