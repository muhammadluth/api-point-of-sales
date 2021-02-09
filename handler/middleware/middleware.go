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
			traceId       = util.CreateUniqID()
			authorization = ctx.Get("Authorization")
		)

		ctx.Set("X-XSS-Protection", "1; mode=block")
		ctx.Set("X-Content-Type-Options", "nosniff")
		ctx.Set("X-Download-Options", "noopen")
		ctx.Set("Strict-Transport-Security", "max-age=5184000")
		ctx.Set("X-Frame-Options", "SAMEORIGIN")
		ctx.Set("X-DNS-Prefetch-Control", "off")

		splitAuthorization := strings.Split(authorization, " ")
		if len(splitAuthorization) != 2 {
			return ctx.Status(fiber.StatusUnauthorized).JSON(model.ResponseHTTP{
				Status:  constant.ERROR,
				Message: "Token Invalid",
			})
		}

		claims, status, err := m.iTokenUsecase.CheckToken(traceId, splitAuthorization[1], ctx)
		if err != nil {
			return ctx.Status(status).JSON(model.ResponseHTTP{
				Status:  constant.ERROR,
				Message: strings.Title(err.Error()),
			})
		}

		splitClaimsSubject := strings.Split(claims.Subject, ":")
		if len(splitClaimsSubject) != 4 {
			return ctx.Status(fiber.StatusUnauthorized).JSON(model.ResponseHTTP{
				Status:  constant.ERROR,
				Message: "Token Invalid",
			})
		}

		if strings.Title(claims.Audience) == strings.Title(constant.ROLE_USER) &&
			strings.ToUpper(ctx.Method()) != strings.ToUpper(constant.METHOD_GET) {
			return ctx.Status(fiber.StatusForbidden).JSON(model.ResponseHTTP{
				Status:  constant.ERROR,
				Message: "Your Role Does Not Have Access",
			})
		}

		ctx.Locals("user_id", splitClaimsSubject[0])
		ctx.Locals("role", claims.Audience)
		return ctx.Next()
	}
}
