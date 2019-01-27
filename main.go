package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var ActiveClients = make(map[ClientConn]int)

type ClientConn struct {
	websocket *websocket.Conn
	clientIP  net.Addr
}

func addClient(cc ClientConn) {
	ActiveClients[cc] = 0
}

func broadcastMessage(messageType int, message []byte) {
	for client, _ := range ActiveClients {
		if err := client.websocket.WriteMessage(messageType, message); err != nil {
			return
		}
	}
}

func deleteClient(cc ClientConn) {
	delete(ActiveClients, cc)
}

func main() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil)

		client := conn.RemoteAddr()
		log.Println(client)
		sockCli := ClientConn{conn, client}
		addClient(sockCli)

		for {
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				deleteClient(sockCli)
				log.Println("bye")
				log.Println(err)
				return
			}
			broadcastMessage(messageType, message)

		}
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.ListenAndServe(":80", nil)
}
