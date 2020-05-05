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

//func GetUser(w http.ResponseWriter, r *http.Request, ps map[string]string) {
//	uc := usecase.GetUseCase()
//	var form forms.RealGetUserForm
//	var user models.User
//	err := json.Unmarshal(data.Body, &form)
//	if err != nil {
//		network.Jsonify(w, forms.ErrorAnswer{
//			Error:   err.Error(),
//			Status:  http.StatusNotAcceptable,
//			Message: "Invalid Json",
//		},  http.StatusNotAcceptable)
//		return
//	}
//	user.Uid = form.Form.Uid
//	code, err := uc.FindUser(&user)
//	if err != nil {
//		network.Jsonify(w, forms.ErrorAnswer{
//			Error:   err.Error(),
//			Status:  code,
//			Message: "something goes wrong while trying to get user",
//		},  code)
//		return
//	}
//	network.Jsonify(w, forms.GetUserAnswer{
//		User:  		user,
//		Status:  	http.StatusOK,
//		Message: 	"successfully get user",
//	}, http.StatusOK)
//}

//func GetProfile(w http.ResponseWriter, r *http.Request, ps map[string]string) {
//	uc := usecase.GetUseCase()
//	var user models.User
//	user.Uid, _ = strconv.Atoi(fmt.Sprintf("%d", r.Context().Value("uid")))
//	code, err := uc.FindUser(&user)
//	if err != nil {
//		network.Jsonify(w, forms.ErrorAnswer{
//			Error:   err.Error(),
//			Status:  code,
//			Message: "something goes wrong while trying to get user",
//		},  code)
//		return
//	}
//	network.Jsonify(w, forms.GetUserAnswer{
//		User:  		user,
//		Status:  	http.StatusOK,
//		Message: 	"successfully get user",
//	}, http.StatusOK)
//}

//func UserLogin(w http.ResponseWriter, r *http.Request, ps map[string]string) {
//	var user forms.RealLoginForm
//	err := json.Unmarshal(data.Body, &user)
//	if err != nil {
//		network.Jsonify(w, forms.ErrorAnswer{
//			Error:   err.Error(),
//			Message: "Invalid Json",
//		},  http.StatusNotAcceptable)
//		// network.GenErrorCode(w, r, "Invalid Json", http.StatusNotAcceptable)
//		return
//	}
//	// form := data.Form
//	//decoder := json.NewDecoder(r.Body)
//	//var form RealForm
//	//// var form middleware.Tokenize
//	//err := decoder.Decode(&form)
//	//if err != nil {
//	//	network.GenErrorCode(w, r, "Invalid Json", http.StatusNotAcceptable)
//	//	return
//	//}
//	//_, err = io.Copy(ioutil.Discard, r.Body)
//	//fmt.Println(r.RemoteAddr)
//	//ok := form.Validate()
//	//if !ok {
//	//	network.ValidationFailed(w, r)
//	//	return
//	//}
//	// toDo move form checker here (from db)
//	uc := usecase.GetUseCase()
//	answer, status, err := uc.LoginUser(user.Form) //form.Form)//form.Form)
//	if err != nil {
//		log.Println("user wasn't logged: db")
//		network.Jsonify(w, forms.ErrorAnswer{
//			Error:   err.Error(),
//			Message: "troubles with db",
//		},  status)
//		return
//	}
//	if status == http.StatusForbidden {
//		log.Println("user wasn't logged: pass")
//		network.Jsonify(w, forms.ErrorAnswer{
//			Error:   "wrong password",
//			Message: "wrong password",
//		},  status)
//		return
//	}
//	if status == http.StatusConflict {
//		log.Println("user wasn't logged: email")
//		network.Jsonify(w, forms.ErrorAnswer{
//			Error:   "no such user: email",
//			Message: "no such user",
//		},  status)
//		return
//	}
//
//
//	muc := usecase.GetSessUseCase()
//	var session models.Session
//	session, err = muc.CreateSession(r.RemoteAddr, r.UserAgent(), answer) // answer - user_id
//	if err != nil {
//		network.Jsonify(w, forms.ErrorAnswer{
//			Error:   err.Error(),
//			Status:  http.StatusNotAcceptable,
//			Message: "create session error",
//		},  http.StatusNotAcceptable)
//		// network.GenErrorCode(w, r, "session creation error", http.StatusInternalServerError)
//		return
//	}
//	go uc.LogSession(session)
//	fmt.Println(session)
//	network.Jsonify(w, forms.LoginAnswer{
//		SessID:  session.SessID,
//		Status:  http.StatusOK,
//		Message: "login succeed",
//	}, status)
//}
