package router

import (
	"flag"
	"fmt"
	"time"

	"api-point-of-sales/app/middleware"
	"api-point-of-sales/handler/authentication"
	"api-point-of-sales/handler/user_management"
	"api-point-of-sales/model"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/muhammadluth/log"
)

type SetupRouter struct {
	timeout                time.Duration
	properties             model.Properties
	iMiddleWare            middleware.Middleware
	iRoleUsecase           user_management.IRoleUsecase
	iUserUsecase           user_management.IUserUsecase
	iRegisterUsecase       authentication.IRegisterUsecase
	iLoginUsecase          authentication.ILoginUsecase
	iForgetPasswordUsecase authentication.IForgetPasswordUsecase
}

func NewSetupRouter(timeout time.Duration, properties model.Properties,
	iMiddleWare middleware.Middleware, iRoleUsecase user_management.IRoleUsecase,
	iUserUsecase user_management.IUserUsecase, iRegisterUsecase authentication.IRegisterUsecase,
	iLoginUsecase authentication.ILoginUsecase,
	iForgetPasswordUsecase authentication.IForgetPasswordUsecase) SetupRouter {
	return SetupRouter{timeout, properties, iMiddleWare, iRoleUsecase,
		iUserUsecase, iRegisterUsecase, iLoginUsecase, iForgetPasswordUsecase}
}

func (r *SetupRouter) Router() {
	addr := flag.String("addr", ":"+r.properties.Port, "TCP address to listen to")
	app := fiber.New()
	app.Use(etag.New())
	app.Use(compress.New())
	app.Use(requestid.New())
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "*",
		AllowHeaders:     "*",
		ExposeHeaders:    "*",
		AllowMethods:     "GET, POST, PUT, DELETE",
	}))

	api := app.Group("/api/v1")

	auth := api.Group("/auth")

	auth.Post("/login", r.iLoginUsecase.Login)
	auth.Post("/register", r.iRegisterUsecase.RegisterUser)
	auth.Post("/forget-password", r.iForgetPasswordUsecase.ForgetPassword)

	role := api.Group("/role")
	role.Get("/", r.iRoleUsecase.GetRoles)
	role.Post("/", r.iMiddleWare.AuthMiddleware(),
		r.iRoleUsecase.CreateRole)

	user := api.Group("/user")
	user.Get("/", r.iMiddleWare.AuthMiddleware(),
		r.iUserUsecase.GetUsers)
	user.Get("/:id", r.iMiddleWare.AuthMiddleware(),
		r.iUserUsecase.GetUserByID)

	// Health Check
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{"message": "Hello, Welcome to My Apps!"})
	})

	log.Event("Listening on " + *addr)
	fmt.Println("Listening on " + *addr)
	fmt.Println("Ready to serve ~")
	log.Fatal(app.Listen(*addr))
}
