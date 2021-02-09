package router

import (
	"flag"
	"fmt"
	"sync"

	"api-point-of-sales/handler"
	"api-point-of-sales/handler/app/authentication"
	"api-point-of-sales/handler/app/user_management"
	"api-point-of-sales/handler/middleware"
	"api-point-of-sales/model"

	"github.com/gofiber/fiber/v2"
	"github.com/muhammadluth/log"
	"github.com/panjf2000/ants/v2"
)

type SetupRouter struct {
	fiberApp               *fiber.App
	poolConnection         *ants.Pool
	properties             model.Properties
	iMiddleWare            middleware.Middleware
	iRoleUsecase           user_management.IRoleUsecase
	iUserUsecase           user_management.IUserUsecase
	iRegisterUsecase       authentication.IRegisterUsecase
	iLoginUsecase          authentication.ILoginUsecase
	iForgetPasswordUsecase authentication.IForgetPasswordUsecase
}

func NewSetupRouter(properties model.Properties,
	iMiddleWare middleware.Middleware,
	iRoleUsecase user_management.IRoleUsecase,
	iUserUsecase user_management.IUserUsecase,
	iRegisterUsecase authentication.IRegisterUsecase,
	iLoginUsecase authentication.ILoginUsecase,
	iForgetPasswordUsecase authentication.IForgetPasswordUsecase) handler.ISetupRouter {
	fiberApp := fiber.New()
	poolConnection, _ := ants.NewPool(int(properties.PoolSize))
	return &SetupRouter{fiberApp, poolConnection, properties, iMiddleWare, iRoleUsecase,
		iUserUsecase, iRegisterUsecase, iLoginUsecase, iForgetPasswordUsecase}
}

func (h *SetupRouter) Router(wg *sync.WaitGroup) {
	addr := flag.String("addr", ":"+h.properties.Port, "TCP address to listen to")

	api := h.fiberApp.Group("/api/v1")

	auth := api.Group("/auth")
	auth.Post("/login", h.iLoginUsecase.Login)
	auth.Post("/register", h.iRegisterUsecase.RegisterUser)
	// auth.Post("/forget-password", h.iRegisterUsecase.RegisterUser)

	role := api.Group("/role")
	role.Get("/", h.iRoleUsecase.GetRoles)
	role.Post("/", h.iMiddleWare.AuthMiddleware(), h.iRoleUsecase.CreateRole)

	user := api.Group("/user")
	user.Get("/", h.iMiddleWare.AuthMiddleware(), h.iUserUsecase.GetUsers)
	user.Get("/:id", h.iMiddleWare.AuthMiddleware(), h.iUserUsecase.GetUserByID)

	// Health Check
	h.fiberApp.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{"message": "Hello, Welcome to My Apps!"})
	})

	log.Event("Listening on " + *addr)
	fmt.Println("Listening on " + *addr)
	fmt.Println("Ready to serve ~")
	log.Fatal(h.fiberApp.Listen(*addr))
}
