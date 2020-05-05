package forms

import "avitocalls/internal/pkg/models"

type CallerAnswer struct {
	BString  string 		`json:"data"`
	Message  string 		`json:"message"`
}

type LongPollAnswer struct {
	Caller   string 		`json:"data"`
	Message  string 		`json:"message"`
}

type ErrorAnswer struct {
	Error  	string			`json:"data"`
	Message string 			`json:"message"`
}


type GetUserAnswer struct {
	User  	models.User		`json:"data"`
	Message string 			`json:"message"`
}

type RegUserAnswer struct {
	UID  	int				`json:"data"`
	Message string 			`json:"message"`
}

type GetUsersAnswer struct {
	Users  	[]models.User	`json:"data"`
	Message string 			`json:"message"`
}

type LoginAnswer struct {
	Users  	[]models.User	`json:"data"`
	UID 	int				`json:"uid"`
	Message string 			`json:"message"`
}

type HistoryCallsAnswer struct {
	Calls 	[]models.Call	`json:"data"`
	Message string			`json:"message"`
}
