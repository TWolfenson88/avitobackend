package server

import (
	"avitocalls/configs/server"
	"avitocalls/internal/pkg/utils"
	"fmt"
	"github.com/gobwas/ws"
	"io"
	"log"
	_ "log"
	"net"
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

func StartTCP() {
	// toDo rewrite normally! то есть вынести настройки и логику, тут оставить только сам запуск
	ln, err := net.Listen("tcp", "0.0.0.0:8100")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tcp is running on 8100")
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		fmt.Println("new conn")
		_, err = ws.Upgrade(conn)
		if err != nil {
			// handle error
		}

		go func() {
			defer conn.Close()

			for {
				fmt.Println("TCP: Opened")
				header, err := ws.ReadHeader(conn)
				if err != nil {
					// handle error
				}

				// смотрим в очередь, а не звонят ли нам

				// a := utils.Rec("calling")
				a := <-utils.ChFirst

				fmt.Printf("get %s, it's ok to do things \n", a)

				payload := make([]byte, header.Length)
				_, err = io.ReadFull(conn, payload)
				if err != nil {
					// handle error
				}
				if header.Masked {
					ws.Cipher(payload, header.Mask, 0)
				}

				// Reset the Masked flag, server frames must not be masked as
				// RFC6455 says.
				header.Masked = false

				if err := ws.WriteHeader(conn, header); err != nil {
					// handle error
				}
				if _, err := conn.Write(payload); err != nil {
					// handle error
				}

				// utils.Send(payload, "answering")
				utils.ChSecond<-payload

				fmt.Println("TCP: Sended")
			}
				//if header.OpCode == ws.OpClose {
				//	return
				//}
			// }
		}()
	}

}
