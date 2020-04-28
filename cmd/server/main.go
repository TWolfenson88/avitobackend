package main

import (
	"avitocalls/internal/app/server"
	_ "log"
)

func main() {
	// log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	go server.StartTCP()
	server.Start()
}
