package network

import (
	"encoding/json"
	"log"
	"net/http"
)

type Message struct {
	Request *http.Request `json:"-"`
	Message string        `json:"message"`
	Status  int           `json:"status"`
}

type LoginAnswer struct {
	SessID     string        `json:"sess_id"`
	Status     int           `json:"status"`
}

func Jsonify(w http.ResponseWriter, object interface{}, status int) {
	output, err := json.Marshal(object)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	w.Header().Set("content-type", "application/json")


	// w.WriteHeader(status)
	_, err = w.Write(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("Sent json")
}


