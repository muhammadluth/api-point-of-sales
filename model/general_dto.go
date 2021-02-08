package model

type (
	ResponseHTTP struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}

	ResponseDropdown struct {
		Label string `json:"label"`
		Value string `json:"value"`
	}

	ResponseSuccessWithoutPagination struct {
		TotalData int         `json:"total_data"`
		Data      interface{} `json:"data"`
	}

	ResponseSuccessWithPagination struct {
		TotalData int         `json:"total_data"`
		TotalPage int         `json:"total_page"`
		Limit     int         `json:"limit"`
		Page      int         `json:"page"`
		Data      interface{} `json:"data"`
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
)

type (
	ParamsUsers struct {
		Limit int `query:"limit"`
		Page  int `query:"page"`
	}
)
