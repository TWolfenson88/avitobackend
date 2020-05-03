package main

import (
	"avitocalls/socket/sock"
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "http service address")

//func serveHome(w http.ResponseWriter, r *http.Request) {
//	log.Println(r.URL)
//	if r.URL.Path != "/" {
//		http.Error(w, "Not found", http.StatusNotFound)
//		return
//	}
//	if r.Method != "GET" {
//		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
//		return
//	}
//	http.ServeFile(w, r, "main.html")
//}

func main() {
	flag.Parse()
	hub := sock.NewHub()
	go hub.Run()
	// http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ServeWs(hub, w, r)
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

//
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	//НЕ ЗАБЫТЬ УБРАТЬ ПОТОМ ВОТ ЕТО ВОТ!!1!
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// serveWs handles websocket requests from the peer.
func ServeWs(hub *sock.Hub, w http.ResponseWriter, r *http.Request) {
	fmt.Println("welcome")
	conn, err := upgrader.Upgrade(w, r, nil)  // do socket conn
	if err != nil {
		log.Println(err)
		return
	}
	//1. Получаем подключение от пользователя.
	//Записываем его indef, который он отправляет первым сообщением при создании сокета
	_, p, err := conn.ReadMessage()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Client connected: ", string(p))
	client := &sock.Client{
		Hub: hub,
		Conn: conn,
		Send: make(chan []byte, 256),
		Indef: string(p),
	}

	//Это должно отправлять список пользователей на клиента. Но оставим это до лучших времен
	//connData, _ := json.Marshal(client.Hub.Clients)
	//
	//errrr := conn.WriteMessage(1, connData)
	//if errrr != nil {
	//	fmt.Println(errrr)
	//}



	client.Hub.Register <- client

	//2. Тута запускаем по 2 рутины на каждого клиента для записи/чтения на сокете
	go client.WritePump()
	go client.ReadPump()
}