package main

import(
	"net/http"
	"encoding/json"
	"fmt"
	"log"
)

type Conection struct {
	From string `json:"from"`
	To   string `json:"to"`
	Cost int    `json:"cost"`
}

type City struct {
	Name string `json:"name"`
}

func main() {
	http.HandleFunc("/", serverRest)
	//http.HandleFunc("/connections/", conectionsFormHandler)
	http.ListenAndServe("localhost:8000", nil)
}

func serverRest(w http.ResponseWriter, r *http.Request) {
	response, err := getJsonResponse()
	if err != nil {
		fmt.Errorf("error: %s", err)
		return
	}

	fmt.Fprintf(w, string(response))
}

func getJsonResponse(){
	conectionForm := make(map[string]string)
}