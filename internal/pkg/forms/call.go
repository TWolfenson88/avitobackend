package forms

import "time"

type CallStartForm struct {
	Caller  	int 			`json:"caller"`
	Answerer  	int 			`json:"answerer"`
	TimeStart  	time.Time		`json:"time"`
}


type CallEndForm struct {
	CallID		int				`json:"call_id"`
	Result		bool			`json:"result"`
	TimeEnd		time.Time		`json:"time"`
}

type StartCallSuccessForm struct {
	CallID		int				`json:"data"`
	Message		string			`json:"message"`
}

type EndCallSuccessForm struct {
	Data		int				`json:"data"`
	Message		string			`json:"message"`
}

type CallHistReqForm struct {
	UID		int					`json:"uid"`
}
