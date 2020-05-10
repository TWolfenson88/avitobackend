package delivery

import (
	"avitocalls/internal/pkg/data"
	"avitocalls/internal/pkg/forms"
	"avitocalls/internal/pkg/models"
	"avitocalls/internal/pkg/network"
	"avitocalls/internal/pkg/user/usecase"
	"encoding/json"
	"log"

	"net/http"
	// "strconv"
)


func FeedUsers(w http.ResponseWriter, r *http.Request, ps map[string]string) {
	uc := usecase.GetUseCase()
	var users []models.User

	users, code, err := uc.InitUsers(users)
	if err != nil {
		network.Jsonify(w, forms.ErrorAnswer{
			Error:   err.Error(),
			Message: "Troubles with initing user feed",
		},  code)
		return
	}

	network.Jsonify(w, forms.GetUsersAnswer{
		Users:  	users,
		Message: 	"successfully get user feed",
	}, http.StatusOK)
}


func RegisterUser(w http.ResponseWriter, r *http.Request, ps map[string]string) {
	uc := usecase.GetUseCase()
	var form models.User
	err := json.Unmarshal(data.Body, &form)
	if err != nil {
		network.Jsonify(w, forms.ErrorAnswer{
			Error:   err.Error(),
			Message: "Invalid Json",
		},  http.StatusNotAcceptable)
		return
	}
	answer, status, err := uc.RegUser(form)  //form.Form)
	if err != nil {
		network.Jsonify(w, forms.ErrorAnswer{
			Error:   err.Error(),
			Message: "user wasn't registered",
		},  http.StatusInternalServerError)
		return
	}
	if status == 409 {
		network.Jsonify(w, forms.RegUserAnswer{
			UID:  		answer,
			Message: 	"user exists",
		}, status)
		return
	}
	network.Jsonify(w, forms.RegUserAnswer{
		UID:  		answer,
		Message: 	"successfully registered user",
	}, status)
}

func LoginUser(w http.ResponseWriter, r *http.Request, ps map[string]string) {
	// toDo move form checker here (from db)
	uc := usecase.GetUseCase()

	var form models.User
	err := json.Unmarshal(data.Body, &form)
	if err != nil {
		network.Jsonify(w, forms.ErrorAnswer{
			Error:   err.Error(),
			Message: "Invalid Json",
		},  http.StatusNotAcceptable)
		return
	}
	uid, status, err := uc.ValidateLogin(form)  // в будущем по уиду надо будет смотреть друзей и их онлайн
	if err != nil {
		log.Println("user wasn't logged cause of db trouble")
		network.Jsonify(w, forms.ErrorAnswer{
			Error:   err.Error(),
			Message: "troubles with db",
		},  status)
		return
	}
	if status == http.StatusForbidden {
		log.Println("user wasn't logged cause of wrong password")
		network.Jsonify(w, forms.ErrorAnswer{
			Error:   "wrong password",
			Message: "wrong password",
		},  status)
		return
	}
	if status == http.StatusConflict {
		log.Println("user wasn't logged cause of wrong name")
		network.Jsonify(w, forms.ErrorAnswer{
			Error:   "no such user: name",
			Message: "no such user",
		},  status)
		return
	}

	var users []models.User
	users, code, err := uc.InitUsers(users)
	if err != nil {
		network.Jsonify(w, forms.ErrorAnswer{
			Error:   err.Error(),
			Message: "Troubles with initing user feed",
		},  code)
		return
	}
	network.Jsonify(w, forms.LoginAnswer{
		Users:  	users,
		UID:		uid,
		Message: 	"successfully get user feed",
	}, http.StatusOK)
}
