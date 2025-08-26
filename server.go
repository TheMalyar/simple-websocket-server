package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// init
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// create variable ws
var clients []*websocket.Conn

func main() {

	// endpoint for connect ws

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		// init upg
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println("Connection sucks", err)
			return
		}

		clients = append(clients, conn)

		// if client send to server
		for {
			// read message from browser
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}

			// print message in term
			fmt.Printf("%s say: %s\n", conn.RemoteAddr(), string(msg))

			for _, client := range clients {
				if err = client.WriteMessage(msgType, msg); err != nil {
					return
				}
			}
		}
	})

	// send html file for browser
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	println("U server run 8080")
	http.ListenAndServe(":8080", nil)

}
