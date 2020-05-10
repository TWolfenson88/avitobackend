package server

import (
	"avitocalls/configs/server"
	"avitocalls/internal/pkg/sock"
	"avitocalls/internal/pkg/sock/delivery"
	"flag"
	"fmt"
	"log"
	_ "log"
	"net/http"
	"strconv"
	"time"
)

func Start() {
	serverSettings := server.GetConfig()
	serve := http.Server{
		Addr:         serverSettings.Ip + ":" + strconv.Itoa(serverSettings.Port),
		Handler:      serverSettings.GetRouter(),
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
	}

	fmt.Println("server is running on " + strconv.Itoa(serverSettings.Port))
	err := serve.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}

var addr = flag.String("addr", ":8080", "http service address")

func SocketStart() {
	flag.Parse()
	hub := sock.NewHub()
	go hub.Run()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		delivery.ServeWs(hub, w, r)
	})
	fmt.Println("Listening socket on 8080")
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
