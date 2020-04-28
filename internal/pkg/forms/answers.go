package forms


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

