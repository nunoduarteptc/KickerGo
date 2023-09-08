package main

import (
	"encoding/json"
	"html/template"
	"kicker/webSockets"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Result struct {
	Id      int    `json:"id"`
	Message string `json:"message"`
}



type Message struct {
	Id      int    `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func postMessage(w http.ResponseWriter, r *http.Request) {
	var result Result
	result.Message = "Recieved message"

	out, err := json.MarshalIndent(result, "", "    ")

	if err != nil {
		log.Println(err)
		return
	}

	var message Message
	_ = json.NewDecoder(r.Body).Decode(&message)

	if err != nil {
		log.Println(err)
		return
	}

	log.Println("string", message.Message)

	// Send message to webSocket

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

	router := mux.NewRouter()

	router.HandleFunc("/home", home).Methods("GET")
	router.HandleFunc("/message", postMessage)
	router.HandleFunc("/api/message", postMessage).Methods("POST")

	log.Println("Starting web server on port 8090")

	http.Handle("/", router)
	http.ListenAndServe(":8090", nil)
}
