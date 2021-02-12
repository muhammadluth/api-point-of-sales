package middleware

import (
	"api-point-of-sales/handler/authentication"
	"api-point-of-sales/model"
	"api-point-of-sales/model/constant"
	"api-point-of-sales/util"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/muhammadluth/log"
	"github.com/panjf2000/ants/v2"
)

type Middleware struct {
	properties     model.Properties
	poolConnection *ants.Pool
	iTokenUsecase  authentication.ITokenUsecase
}

func NewMiddleware(properties model.Properties, iTokenUsecase authentication.ITokenUsecase) Middleware {
	poolConnection, _ := ants.NewPool(int(properties.PoolSize), ants.WithPreAlloc(true))
	return Middleware{properties, poolConnection, iTokenUsecase}
}

func (m *Middleware) AuthMiddleware() func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		var (
			uniqId        = util.CreateUniqID()
			authorization = ctx.Get("Authorization")
		)

		m.poolConnection.Submit(func() {
			log.Message(
				uniqId,
				"IN",
				"GO-FIBER",
				"",
				"URL",
				ctx.OriginalURL(),
				"",
				"REQUEST",
				string(ctx.Request().Body()))
		})

		splitAuthorization := strings.Split(authorization, " ")
		if len(splitAuthorization) != 2 {
			return ctx.Status(fiber.StatusUnauthorized).JSON(model.ResponseHTTP{
				Status:  constant.ERROR,
				Message: "Token Invalid",
			})
		}

		claims, status, err := m.iTokenUsecase.CheckToken(uniqId, splitAuthorization[1], ctx)
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

		ctx.Set("X-XSS-Protection", "1; mode=block")
		ctx.Set("X-Content-Type-Options", "nosniff")
		ctx.Set("X-Download-Options", "noopen")
		ctx.Set("Strict-Transport-Security", "max-age=5184000")
		ctx.Set("X-Frame-Options", "SAMEORIGIN")
		ctx.Set("X-DNS-Prefetch-Control", "off")

		ctx.Locals("uniqId", uniqId)
		ctx.Locals("user_id", splitClaimsSubject[0])
		ctx.Locals("role", claims.Audience)

		m.poolConnection.Submit(func() {
			log.Message(
				uniqId,
				"OUT",
				"GO-FIBER",
				"",
				"URL",
				ctx.OriginalURL(),
				"",
				"RESPONSE",
				string(ctx.Response().Body()))
		})
		return ctx.Next()
	}
}
