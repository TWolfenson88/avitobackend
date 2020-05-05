package forms

import "avitocalls/internal/pkg/models"

type CallerAnswer struct {
	BString  string 		`json:"data"`
	Message  string 		`json:"message"`
	// Status   int 			`json:"status"`
}

type LongPollAnswer struct {
	Caller   string 		`json:"data"`
	Message  string 		`json:"message"`
	// Status   int 			`json:"status"`
}

type ErrorAnswer struct {
	Error  	string			`json:"data"`
	Message string 			`json:"message"`
	// Status  int 			`json:"status"`
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

