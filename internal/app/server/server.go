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
	// err := server.ListenAndServeTLS("./configs/ssl-bundle/bundle.crt", "./configs/ssl-bundle/private.key.pem")
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
	// http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		delivery.ServeWs(hub, w, r)
	})
	fmt.Println("Listening socket on 8080")
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}










//func StartTCP() {
//	go func() {
//		time.Sleep(7*time.Second)
//		utils.ChFirst <- "abracadabra"
//		fmt.Println("HAVE SENT")
//	}()
//
//	// toDo rewrite normally! то есть вынести настройки и логику, тут оставить только сам запуск
//	ln, err := net.Listen("tcp", "0.0.0.0:8100")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println("tcp is running on 8100")
//	for {
//		conn, err := ln.Accept()
//		if err != nil {
//			// handle error
//		}
//		fmt.Println("new conn")
//		_, err = ws.Upgrade(conn)
//		if err != nil {
//			// handle error
//		}
//		// rofl := 0
//		go func() {
//			defer conn.Close()
//			//header, err := ws.ReadHeader(conn)
//			//if err != nil {
//			//	// handle error
//			//}
//			for {
//				// 1 Кирилл коннектится
//				fmt.Println("TCP: Opened")
//				header, _ := ws.ReadHeader(conn)
//				if len(utils.ChKirillOnline) == 0 {
//					fmt.Println("TCP: Setting KIRILL to online")
//					// 2 ставим Кириллу онлайн путём заполнения канала
//					utils.ChKirillOnline <- "ONLINE"
//				}
//				// 3 ждем, пока от Андрея придет запрос на звонок
//				andrey := <- utils.ChAndrey
//				fmt.Println("TCP: got ", andrey)
//				// 7 вычитываем все, что пришло от Кирилла (сдп-шник)
//				payload := make([]byte, header.Length)
//				_, err = io.ReadFull(conn, payload)
//				if err != nil { panic(err) }
//				if header.Masked { ws.Cipher(payload, header.Mask, 0) }
//				header.Masked = false
//
//				if err := ws.WriteHeader(conn, header); err != nil { panic(err) }
//				// 8 Отправляем сдп обратно Кириллу (а чего бы и нет) даже если он уже закрыл страницу
//				if _, err := conn.Write(payload); err != nil {
//					// handle error
//				}
//				// 9 Кладем в канал Кирилла СДП (на всякий случай проверяем, а не лежит ли что-то уже в канале,
//				// что маловероятно, но проверить стоит)
//				if len(utils.ChKirill) == 0 {
//					utils.ChKirill <- payload
//					fmt.Println("TCP: put payload to channel")
//				} else {
//					fmt.Println("TCP: Kirill's channel not empty")
//				}
//				// 10 убираем Кирилла из онлайна, чтобы нельзя было снова выполнить п. 5
//				if len(utils.ChKirillOnline) == 1 {
//					fmt.Println("TCP: Setting KIRILL to offline")
//					_ = <-utils.ChKirillOnline
//				}
//
//				// utils.ChSecond<-payload
//
//				fmt.Println("TCP: Closing")
//				return
//
//				//fmt.Println("HeaderCode", header.OpCode)
//				//if header.OpCode == ws.OpClose {
//				//	rofl += 1
//				//} else {
//				//	rofl = 0
//				//}
//				//if rofl >= 3 {
//				//	fmt.Println("Rofl closed me")
//				//	return
//				//}
//				//if header.OpCode == ws.OpClose {
//				//	// тут убираем челика из онлайна
//				//	if len(utils.ChKirillOnline) == 1 {
//				//		fmt.Println("TCP: Setting KIRILL to offline")
//				//
//				//		payload := make([]byte, header.Length)
//				//		_, err = io.ReadFull(conn, payload)
//				//		if err != nil { panic(err) }
//				//		if header.Masked { ws.Cipher(payload, header.Mask, 0) }
//				//		header.Masked = false
//				//		if err := ws.WriteHeader(conn, header); err != nil { panic(err) }
//				//		if _, err := conn.Write(payload); err != nil {
//				//			// handle error
//				//		}
//				//
//				//		fmt.Println(rofl)
//				//		_ = <- utils.ChKirillOnline
//				//	}
//				//	fmt.Println("TCP: Closing conn")
//				//	return
//				//}
//			}
//
//			// }
//		}()
//	}
//
//}
