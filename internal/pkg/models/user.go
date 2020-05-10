package models

type User struct {
	Uid 			int 	  `json:"uid"`
	Name 			string    `json:"name"`
	//Email 			string    `json:"email, 	omitempty"`
	Password        string    `json:"password, 	omitempty"`
	//Identify		string	  `json:"ident, 	omitempty"`
	StatusOnline	bool	  `json:"online, 	omitempty"`
}

