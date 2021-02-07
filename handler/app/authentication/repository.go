package authentication

import "api-point-of-sales/model"

type IAuthenticationRepo interface {
	GetUserByUsername(traceId, account string) (*model.TableUsers, error)
	GetUserByEmail(traceId, account string) (*model.TableUsers, error)
}
