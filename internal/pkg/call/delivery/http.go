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
	// 4 Андрей заходит
	defer func() {
		fmt.Println("HTTP: Andrey closed connection. Setting him off")
		if len(utils.ChAndrey) == 1 {
			_ = <- utils.ChAndrey
		}
	}()
	// 5 Андрей проверяет в онлайне ли Кирилл (положил ли он что-то в канал, см п.2)
	if len(utils.ChKirillOnline) == 0 {
		fmt.Println("HTTP: Kirill if offline")
		// 5.2 Если канал Кирилла пусть, что значит, что Кирилл не зашел
		network.Jsonify(w, "KIRILL OFFLINE", http.StatusOK)
		return
	}
	// 6 Андрей кладет данные в канал, который слушает Кирилл (см п. 3)
	if len(utils.ChAndrey) == 0 {
		utils.ChAndrey<-"Andrey wants to call Kirill"
		fmt.Println("HTTP: Put to ChAndrey")
	} else {
		// 6.1 Такая ситуация маловероятна, но я её обработал
		fmt.Println("HTTP: ChAndrey not empty")
		network.Jsonify(w, "Another Andrey is online?? SHIT! Mb needs to restart server??",
			http.StatusInternalServerError)
		return
	}
	// 6.2 Андрей ждет 100 мс, чтобы Кирилл точно успел отправить Андрею сдп
	time.Sleep(100 * time.Millisecond)
	// 11 Кирилл должен был усппеть записать сдп в свой канал, но проверим этот факт, на всякий случай
	if len(utils.ChKirill) == 0 {
		fmt.Println("HTTP: ChKirill is empty")
		network.Jsonify(w, "Somebody red from Kirill. SHIT!", http.StatusInternalServerError)
		return
	}
	// 11.1 вычитываем сдп из канала Кирилла
	res := <-utils.ChKirill
	fmt.Printf("HTTP: %s", res)

	// 12 оборачиваем сдп в джсон
	var sdp forms.SDP
	err := json.Unmarshal(res, &sdp)
	if err != nil {
		network.Jsonify(w, forms.ErrorAnswer{
			Error:   err.Error(),
			Message: "Shit with answer format: try again",
		}, http.StatusInternalServerError)
		return
	}
	// 13 отправляем Андрею
	network.Jsonify(w, sdp, http.StatusOK)
	// 14 У Андрюхи есть сдп Кирилла. ПРОФИТ!
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
