package forms

type RealGetUserForm struct {
	Form GetUserForm 		`json:"data"`
}

type GetUserForm struct {
	Uid 	int 			`json:"uid"`
}

type RealRegUserForm struct {
	Form RegUserForm 		`json:"data"`
}

type RegUserForm struct {
	Name 		string 		`json:"name"`
	Password 	string 		`json:"password"`
}
