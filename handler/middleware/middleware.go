package middleware

import (
	"api-point-of-sales/handler/app/authentication"
	"api-point-of-sales/model"
	"api-point-of-sales/model/constant"
	"api-point-of-sales/util"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type Middleware struct {
	iTokenUsecase authentication.ITokenUsecase
}

func NewMiddleware(iTokenUsecase authentication.ITokenUsecase) Middleware {
	return Middleware{iTokenUsecase}
}

func (m *Middleware) AuthMiddleware() func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		var (
			traceId       = util.CreateTraceID()
			authorization = ctx.Get("Authorization")
		)
		splitAuthorization := strings.Split(authorization, " ")
		if len(splitAuthorization) != 2 {
			return ctx.Status(fiber.StatusUnauthorized).JSON(model.ResponseHTTP{
				Status:  constant.ERROR,
				Message: "Token invalid",
			})
		}

		claims, status, err := m.iTokenUsecase.CheckToken(traceId, splitAuthorization[1], ctx)
		if err != nil {
			return ctx.Status(status).JSON(model.ResponseHTTP{
				Status:  constant.ERROR,
				Message: strings.Title(err.Error()),
			})
		}

		ctx.Set("X-XSS-Protection", "1; mode=block")
		ctx.Set("X-Content-Type-Options", "nosniff")
		ctx.Set("X-Download-Options", "noopen")
		ctx.Set("Strict-Transport-Security", "max-age=5184000")
		ctx.Set("X-Frame-Options", "SAMEORIGIN")
		ctx.Set("X-DNS-Prefetch-Control", "off")
		ctx.Locals("id", claims.Id)
		return ctx.Next()
	}
}
