package forms

import "avitocalls/internal/pkg/models"

type RealLoginForm struct {
	Form LoginForm `json:"data"`
}

type RealRegForm struct {
	Form models.User `json:"data"`
}

type LoginForm struct {
	Name	   	string      `json:"name"`
	Password 	string 		`json:"password"`
}

