package webSockets

import (
	"log"
	"net/http"
	"github.com/gorilla/websocket" // Old websocket need to improve
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:1024,
	WriteBufferSize:1024,
}

func reader(conn *websocket.Conn) {
	for {
		messageType, message ,err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		messageToSent := []byte("Hello Back")
		log.Println(string(message))
		if err := conn.WriteMessage(messageType, messageToSent); err != nil {
			log.Println(err)
			return
		}
	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {return true}

	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
	}

	log.Printf("Client Successfully connected")

	reader(ws)
}

func Init() {
	http.HandleFunc("/ws", wsEndpoint)

	log.Println("Starting web server on port 8091")

	http.ListenAndServe(":8091", nil)
}

