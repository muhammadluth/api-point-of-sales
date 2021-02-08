package model

type (
	RequestCreateRole struct {
		RoleName    string `json:"role_name" xml:"role_name" form:"role_name"`
		Description string `json:"description" xml:"description" form:"description"`
	}

	RequestCreateUser struct {
		FirstName       string `json:"first_name" xml:"first_name" form:"first_name"`
		LastName        string `json:"last_name" xml:"last_name" form:"last_name"`
		Username        string `json:"username" xml:"username" form:"username"`
		Email           string `json:"email" xml:"email" form:"email"`
		PhoneNumber     string `json:"phone_number" xml:"phone_number" form:"phone_number"`
		Password        string `json:"password" xml:"password" form:"password"`
		ConfirmPassword string `json:"confirm_password" xml:"confirm_password" form:"confirm_password"`
		RoleID          string `json:"role_id" xml:"role_id" form:"role_id"`
	}
)

type (
	ResponseGetUsers struct {
		UserID      string `json:"user_id"`
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		Username    string `json:"username"`
		Email       string `json:"email"`
		PhoneNumber string `json:"phone_number"`
		Role        string `json:"role"`
	}
)
