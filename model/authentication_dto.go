package model

type (
	RequestLogin struct {
		Account  string `json:"account" xml:"account" form:"account"`
		Password string `json:"password" xml:"password" form:"password"`
	}
)

type (
	ResponseLogin struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Role      string `json:"role"`
	}
)
