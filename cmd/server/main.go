package main

import (
	"avitocalls/internal/app/server"
	_ "log"
)

func main() {
	go server.SocketStart()
	server.Start()

}
