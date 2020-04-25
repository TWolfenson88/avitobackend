package delivery

import (
	"avitocalls/internal/pkg/call/usecase"
	"avitocalls/internal/pkg/data"
	"avitocalls/internal/pkg/forms"
	"avitocalls/internal/pkg/models"
	"avitocalls/internal/pkg/network"
	"avitocalls/internal/pkg/stun"
	"encoding/json"
	"net/http"
)

func CallUser(w http.ResponseWriter, r *http.Request, ps map[string]string) {
	uc := usecase.GetUseCase()
	var call models.Call
	err := json.Unmarshal(data.Body, &call)

	// connect to stun server using
	_, err = stun.ConnectToSTUN()
	if err != nil {
		network.Jsonify(w, forms.ErrorAnswer{
			Error:   err.Error(),
			Status:  http.StatusInternalServerError,
			Message: "Connecting to stun-server error",
		}, http.StatusInternalServerError)
		return
	}

	call, code, err := uc.GetReceiverObject(call)

	// making call func

	if err != nil {
		network.Jsonify(w, forms.ErrorAnswer{
			Error:   err.Error(),
			Status:  http.StatusInternalServerError,
			Message: "Error",
		}, http.StatusInternalServerError)
		return
	}

	var stringg = "here i am"
	switch code {
	case 200:
		network.Jsonify(w, forms.CallerAnswer{
			BString: stringg,
			Status:  http.StatusOK,
			Message: "successfully made call",
		}, http.StatusOK)
	case http.StatusExpectationFailed:
		network.Jsonify(w, forms.CallerAnswer{
			BString: stringg,
			Status:  http.StatusExpectationFailed,
			Message: "user do not accepted call",
		}, http.StatusOK)
	}
}
