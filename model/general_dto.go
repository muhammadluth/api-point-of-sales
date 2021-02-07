package model

type (
	ResponseHTTP struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}

	DataUser struct {
		UserID      string   `json:"user_id"`
		FirstName   string   `json:"first_name"`
		LastName    string   `json:"last_name"`
		Username    string   `json:"username"`
		Email       string   `json:"email"`
		PhoneNumber string   `json:"phone_number"`
		Password    string   `json:"password"`
		Role        DataRole `json:"role_id"`
	}

	DataRole struct {
		RoleName    string `json:"role_name"`
		Description string `json:"description"`
	}

	PayloadToken struct {
		Jti string `json:"jti"`
		Iss string `json:"iss"`
		Sub string `json:"sub"`
		Aud string `json:"aud"`
		Exp int64  `json:"exp"`
		Nbf int64  `json:"nbf"`
		Iat int64  `json:"iat"`
	}
)
