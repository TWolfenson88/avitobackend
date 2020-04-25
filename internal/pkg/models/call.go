package models

type Call struct {
	CallerID 		string    `json:"caller"`
	ReceiverID		string    `json:"receiver"`
	BString 		string 	  `json:"b64"`
	// что-то ещё?
}