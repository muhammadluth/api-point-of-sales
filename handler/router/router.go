package router

import (
	"flag"
	"fmt"
	"sync"

	"api-point-of-sales/handler"
	"api-point-of-sales/handler/app/authentication"
	"api-point-of-sales/handler/app/user_management"
	"api-point-of-sales/handler/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/muhammadluth/log"
	"github.com/panjf2000/ants/v2"
)

type SetupRouter struct {
	fiberApp         *fiber.App
	poolConnection   *ants.Pool
	iMiddleWare      middleware.Middleware
	iRoleUsecase     user_management.IRoleUsecase
	iUserUsecase     user_management.IUserUsecase
	iRegisterUsecase authentication.IRegisterUsecase
	iLoginUsecase    authentication.ILoginUsecase
}

func NewSetupRouter(poolSize int,
	iMiddleWare middleware.Middleware,
	iRoleUsecase user_management.IRoleUsecase,
	iUserUsecase user_management.IUserUsecase,
	iRegisterUsecase authentication.IRegisterUsecase,
	iLoginUsecase authentication.ILoginUsecase) handler.ISetupRouter {
	fiberApp := fiber.New()
	poolConnection, _ := ants.NewPool(poolSize)
	return &SetupRouter{fiberApp, poolConnection, iMiddleWare, iRoleUsecase, iUserUsecase,
		iRegisterUsecase, iLoginUsecase}
}

func (h *SetupRouter) Router(wg *sync.WaitGroup) {
	addr := flag.String("addr", ":"+"8081", "TCP address to listen to")

	api := h.fiberApp.Group("/api/v1", logger.New())

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", h.iLoginUsecase.Login)
	auth.Post("/register", h.iRegisterUsecase.RegisterUser)

	// Role
	role := api.Group("/role")
	role.Get("/", h.iRoleUsecase.GetRoles)
	role.Post("/", h.iMiddleWare.AuthMiddleware(), h.iRoleUsecase.CreateRole)

	// User
	user := api.Group("/user")
	user.Get("/", h.iMiddleWare.AuthMiddleware(), h.iUserUsecase.GetUsers)

	// Health Check
	h.fiberApp.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{"message": "Hello, Welcome to My Apps!"})
	})

	log.Event("Listening on " + *addr)
	fmt.Println("Listening on " + *addr)
	fmt.Println("Ready to serve ~")
	log.Fatal(h.fiberApp.Listen(*addr))
}
