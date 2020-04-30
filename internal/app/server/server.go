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
	go func() {
		time.Sleep(7*time.Second)
		utils.ChFirst <- "abracadabra"
		fmt.Println("HAVE SENT")
	}()

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
		// rofl := 0
		go func() {
			defer conn.Close()
			//header, err := ws.ReadHeader(conn)
			//if err != nil {
			//	// handle error
			//}
			for {
				fmt.Println("TCP: Opened")
				header, _ := ws.ReadHeader(conn)
				if len(utils.ChKirillOnline) == 0 {
					fmt.Println("TCP: Setting KIRILL to online")
					utils.ChKirillOnline <- "ONLINE"
				}

				andrey := <- utils.ChAndrey
				fmt.Println("TCP: got ", andrey)


				// ждем, пока в канал что-то не придет
				// _ = <-utils.ChFirst

				payload := make([]byte, header.Length)
				_, err = io.ReadFull(conn, payload)
				if err != nil { panic(err) }
				if header.Masked { ws.Cipher(payload, header.Mask, 0) }
				header.Masked = false

				if err := ws.WriteHeader(conn, header); err != nil { panic(err) }

				// fmt.Println(len(utils.ChKirillOnline))

				if _, err := conn.Write(payload); err != nil {
					// handle error
				}


				if len(utils.ChKirill) == 0 {
					utils.ChKirill <- payload
					fmt.Println("TCP: put payload to channel")
				} else {
					fmt.Println("TCP: Kirill's channel not empty")
				}
				if len(utils.ChKirillOnline) == 1 {
					fmt.Println("TCP: Setting KIRILL to offline")
					_ = <-utils.ChKirillOnline
				}


				// utils.ChSecond<-payload

				fmt.Println("TCP: Closing")
				return

				//fmt.Println("HeaderCode", header.OpCode)
				//if header.OpCode == ws.OpClose {
				//	rofl += 1
				//} else {
				//	rofl = 0
				//}
				//if rofl >= 3 {
				//	fmt.Println("Rofl closed me")
				//	return
				//}
				//if header.OpCode == ws.OpClose {
				//	// тут убираем челика из онлайна
				//	if len(utils.ChKirillOnline) == 1 {
				//		fmt.Println("TCP: Setting KIRILL to offline")
				//
				//		payload := make([]byte, header.Length)
				//		_, err = io.ReadFull(conn, payload)
				//		if err != nil { panic(err) }
				//		if header.Masked { ws.Cipher(payload, header.Mask, 0) }
				//		header.Masked = false
				//		if err := ws.WriteHeader(conn, header); err != nil { panic(err) }
				//		if _, err := conn.Write(payload); err != nil {
				//			// handle error
				//		}
				//
				//		fmt.Println(rofl)
				//		_ = <- utils.ChKirillOnline
				//	}
				//	fmt.Println("TCP: Closing conn")
				//	return
				//}
			}

			// }
		}()
	}

}
