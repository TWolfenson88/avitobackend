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
	//1. Получаем подключение от пользователя.
	//Записываем его indef, который он отправляет первым сообщением при создании сокета
	_, p, err := conn.ReadMessage()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Client connected: ", string(p))


	// toDO write p to ONLINE
	if p != nil {
		//go func() {
			uc := usecase.GetUseCase()
			uc.SetOnline(string(p))
		//}()
	}



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