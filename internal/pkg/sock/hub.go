package sock

import (
	"encoding/json"
	"fmt"
	"log"
)


// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	Clients map[*Client]bool

	// Inbound messages from the clients.
	Broadcast chan []byte

	// Register requests from the clients.
	Register chan *Client

	// Unregister requests from clients.
	Unregister chan *Client
}
//Тут надо понять как друг другу передать объект
type Msgg struct {
	Receiver 	string `json:"receiver"`
	Obj 		string `json:"obj"`
}


func NewHub() *Hub {
	return &Hub{
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	//3. Тут у нас бесконечный цикл Хаба, который держит на себе все соединения и проверяет кто куда, раздаёт сообщения от клиента к клиенту
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
		case client := <-h.Unregister:
			h.Clients[client] = true
		/*
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		*/
		case message := <-h.Broadcast:
			for client := range h.Clients {
				fmt.Println("message in broadcast: ", string(message))

				//5. Вот тут происходит парс ЖЫСОНа и передача на клиента б64 строки (obj)
				var msg = Msgg{}

				errr := json.Unmarshal(message, &msg)
				if errr != nil {
					log.Fatal("error unmarshall")
				}

				fmt.Println(h.Clients)
				// {"type":"userlist", "userlist":"....."}
				if msg.Receiver == client.Indef {
					select {
					case client.Send <- []byte(msg.Obj):
					default:
						fmt.Println("there?")
						close(client.Send)
						delete(h.Clients, client)
					}
				}
			}
		}
	}
}

