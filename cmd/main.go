package main

import(
	"net/http"
	"encoding/json"
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
	http.HandleFunc("/connections/", conectionsFormHandler)
	http.ListenAndServe(":8000", nil)
}

func conectionsFormHandler(w http.ResponseWriter, r *http.Request) {
	conectionForm := Conection{}
	
	// FIXME: should check errors!
	req, err := json.NewDecoder(r.Body).Decode(&conectionForm)
	if err != nil{
		log.Fatal("Error while reading r.Body: ", err)
}
}