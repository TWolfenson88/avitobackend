package delivery

import (
	"avitocalls/internal/pkg/sock"
	"avitocalls/internal/pkg/user/usecase"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

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
	_, p, err := conn.ReadMessage()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Client connected: ", string(p))

	// toDO write p to ONLINE
	if p != nil {
		go func() {
			uc := usecase.GetUseCase()
			err = uc.SetOnline(string(p))
			if err != nil {
				fmt.Println("error with setting online", p, err)
			}
		}()
	}

	client := &sock.Client{
		Hub: hub,
		Conn: conn,
		Send: make(chan []byte, 256),
		Indef: string(p),
	}

	client.Hub.Register <- client

	go client.WritePump()
	go client.ReadPump()
}
