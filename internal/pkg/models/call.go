package models

import "time"

//type Call struct {
//	CallerID 		string    `json:"caller"`
//	ReceiverID		string    `json:"receiver"`
//	BString 		string 	  `json:"b64"`
//	// что-то ещё?
//}


type Call struct {
	CallID			int 			`json:"call_id"`
	CallerID		int 			`json:"caller_id"`
	CallerName		string 			`json:"call_name"`
	AnswererID		int 			`json:"answerer_id"`
	AnswererName	string 			`json:"answerer_name"`
	StartTime		time.Time		`json:"start_time"`
	Duration		int				`json:"duration_sec"`
	Result			bool			`json:"result"`
}