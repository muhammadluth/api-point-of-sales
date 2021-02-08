package constant

const (
	CONNECT_DB = "postgres://{:db_user}:{:db_password}@{:db_host}:{:db_port}/{:db_name}"
)

const (
	METHOD_GET    = "GET"
	METHOD_POST   = "POST"
	METHOD_PUT    = "PUT"
	METHOD_DELETE = "DELETE"
)

const (
	ROLE_ADMIN = "Admin"
	ROLE_USER  = "User"
)
