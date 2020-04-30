package delivery

import (
	"avitocalls/internal/pkg/forms"
	"avitocalls/internal/pkg/network"
	"avitocalls/internal/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)


func WaitForCall(w http.ResponseWriter, r *http.Request, ps map[string]string) {
	// пока это просто заглушка, в будущем тут надо проверять, не звонят ли юзеру в данный момент
	if true {
		network.Jsonify(w, forms.LongPollAnswer{
			Caller:  "NoName",
			Message: "nobody wants to call you",
		}, http.StatusAccepted)
	} else {
		network.Jsonify(w, forms.LongPollAnswer{
			Caller:  "Name",
			Message: "user wants to make a call",
		}, http.StatusOK)
	}
	return
}

func CallUser(w http.ResponseWriter, r *http.Request, ps map[string]string) {
	defer func() {
		fmt.Println("HTTP: Andrey occasionally closed connection. Setting him off")
		if len(utils.ChAndrey) == 1 {
			_ = <- utils.ChAndrey
		}
	}()

	if len(utils.ChKirillOnline) == 0 {
		fmt.Println("HTTP: Kirill if offline")
		network.Jsonify(w, "KIRILL OFFLINE", http.StatusOK)
		return
	}

	if len(utils.ChAndrey) == 0 {
		utils.ChAndrey<-"Andrey wants to call Kirill"
		fmt.Println("HTTP: Put to ChAndrey")
	} else {
		fmt.Println("HTTP: ChAndrey not empty")
		network.Jsonify(w, "Another Andrey is online?? SHIT! Mb needs to restart server??",
			http.StatusInternalServerError)
		return
	}
	time.Sleep(100 * time.Millisecond)
	if len(utils.ChKirill) == 0 {
		fmt.Println("HTTP: ChKirill is empty")
		network.Jsonify(w, "Somebody red from Kirill. SHIT!", http.StatusInternalServerError)
		return
	}
	res := <-utils.ChKirill
	fmt.Printf("HTTP: %s", res)

	var sdp forms.SDP
	err := json.Unmarshal(res, &sdp)
	if err != nil {
		network.Jsonify(w, forms.ErrorAnswer{
			Error:   err.Error(),
			Message: "Shit with answer format: try again",
		}, http.StatusInternalServerError)
		return
	}
	network.Jsonify(w, sdp, http.StatusOK)
	return

	// by RabbitMQ
	//utils.Send([]byte("so call me maybe"), "calling")
	//
	//res := utils.Rec("answering")
	//var sdp forms.SDP
	//err := json.Unmarshal(res, &sdp)
	//if err != nil {
	//	network.Jsonify(w, forms.ErrorAnswer{
	//		Error:   err.Error(),
	//		Message: "Shit with answer format: try again",
	//	}, http.StatusInternalServerError)
	//}
	//network.Jsonify(w, sdp, http.StatusOK)


	// by accident :)
	// uc := usecase.GetUseCase()
	// var call models.Call
	// err := json.Unmarshal(data.Body, &call)

	// connect to stun server using
	//_, err = stun.ConnectToSTUN()
	//if err != nil {
	//	network.Jsonify(w, forms.ErrorAnswer{
	//		Error:   err.Error(),
	//		Message: "Connecting to stun-server error",
	//	}, http.StatusInternalServerError)
	//	return
	//}
	//
	//call, code, err := uc.GetReceiverObject(call)
	//
	//// making call func
	//
	//if err != nil {
	//	network.Jsonify(w, forms.ErrorAnswer{
	//		Error:   err.Error(),
	//		Message: "Error",
	//	}, http.StatusInternalServerError)
	//	return
	//}
	//
	//var stringg = "here i am"
	//switch code {
	//case http.StatusOK:
	//	network.Jsonify(w, forms.CallerAnswer{
	//		BString: stringg,
	//		Message: "successfully made call",
	//	}, http.StatusOK)
	//case http.StatusExpectationFailed:
	//	network.Jsonify(w, forms.CallerAnswer{
	//		BString: stringg,
	//		Message: "user do not accepted call",
	//	}, http.StatusOK)
	//}
	//return
}
