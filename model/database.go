package model

//SCHEMA AUTHENTICATION
type (
	TableUsers struct {
		tableName   struct{} `pg:"authentication.users"`
		UserID      string   `pg:",pk"`
		FirstName   string
		LastName    string
		Username    string
		Email       string
		PhoneNumber string
		Password    string
		RoleID      string
		Roles       *TableRoles `pg:"rel:has-one"`
	}

	TableRoles struct {
		tableName   struct{} `pg:"authentication.roles"`
		RoleID      string   `pg:",pk"`
		RoleName    string
		Description string
	}
)
