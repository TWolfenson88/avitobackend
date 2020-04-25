package server

import (
	"avitocalls/configs/server"
	"fmt"
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
	// err := server.ListenAndServeTLS("./configs/ssl-bundle/bundle.crt", "./configs/ssl-bundle/private.key.pem")
	err := serve.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
