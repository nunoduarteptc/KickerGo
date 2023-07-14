package main

import (
	"html/template"
	"encoding/json"
	"log"
	"net/http"
	"kicker/webSockets"
)

type Result struct {
	Id int `json:"id"`
	Message string `json:"message"`
}


func postMessage(w http.ResponseWriter, r *http.Request) {
	var result Result
	result.Message = "Recieved message"

	out, err := json.MarshalIndent(result, "", "    ")

	if (err != nil) {
		log.Println(err)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.Write(out)
}

func home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")

	if err != nil {
		log.Println(err)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err)
		return
	}
}

func main() {
	go webSockets.Init()
	http.HandleFunc("/home", home)
	http.HandleFunc("/message", postMessage)

	log.Println("Starting web server on port 8090")
	http.ListenAndServe(":8090", nil)
}


