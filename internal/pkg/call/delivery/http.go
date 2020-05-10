package delivery

import (
	"avitocalls/internal/pkg/call/usecase"
	"avitocalls/internal/pkg/data"
	"avitocalls/internal/pkg/forms"
	"avitocalls/internal/pkg/network"
	"encoding/json"
	"net/http"
)


func StartCall(w http.ResponseWriter, r *http.Request, ps map[string]string) {
	uc := usecase.GetUseCase()
	var form forms.CallStartForm
	err := json.Unmarshal(data.Body, &form)
	if err != nil {
		network.Jsonify(w, forms.ErrorAnswer{
			Error:   err.Error(),
			Message: "Invalid Json",
		},  http.StatusNotAcceptable)
		return
	}

	callid, err := uc.SaveCallStarting(form)
	if err != nil {
		network.Jsonify(w, forms.ErrorAnswer{
			Error:   err.Error(),
			Message: "start call info wasn't saved",
		}, http.StatusInternalServerError)
		return
	}
	network.Jsonify(w, forms.StartCallSuccessForm{
		CallID:  		callid,
		Message: 		"start call info saved",
	}, http.StatusOK)
}

func EndCall(w http.ResponseWriter, r *http.Request, ps map[string]string) {
	uc := usecase.GetUseCase()
	var form forms.CallEndForm
	err := json.Unmarshal(data.Body, &form)
	if err != nil {
		network.Jsonify(w, forms.ErrorAnswer{
			Error:   err.Error(),
			Message: "Invalid Json",
		},  http.StatusNotAcceptable)
		return
	}

	_, err = uc.SaveCallEnding(form)
	if err != nil {
		network.Jsonify(w, forms.ErrorAnswer{
			Error:   err.Error(),
			Message: "end call info wasn't saved",
		}, http.StatusInternalServerError)
		return
	}
	network.Jsonify(w, forms.EndCallSuccessForm{
		Data:  			1,
		Message: 		"end call info saved",
	}, http.StatusOK)
}

func GetHistory(w http.ResponseWriter, r *http.Request, ps map[string]string) {
	uc := usecase.GetUseCase()
	var form forms.CallHistReqForm
	err := json.Unmarshal(data.Body, &form)
	if err != nil {
		network.Jsonify(w, forms.ErrorAnswer{
			Error:   err.Error(),
			Message: "Invalid Json",
		},  http.StatusNotAcceptable)
		return
	}

	calls, err := uc.GetUserCallHistory(form)
	if err != nil {
		network.Jsonify(w, forms.ErrorAnswer{
			Error:   err.Error(),
			Message: "can't get history of calls",
		}, http.StatusInternalServerError)
		return
	}
	network.Jsonify(w, forms.HistoryCallsAnswer{
		Calls:  		calls,
		Message: 		"successfully get call history",
	}, http.StatusOK)
}


