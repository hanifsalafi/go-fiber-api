package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"go-fiber-api/app/module/user"
	"go-fiber-api/config/config"
)

type Router struct {
	App fiber.Router
	Cfg *config.Config

	UserRouter *user.UserRouter
}

func NewRouter(
	fiber *fiber.App,
	cfg *config.Config,

	userRouter *user.UserRouter,
) *Router {
	return &Router{
		App:        fiber,
		Cfg:        cfg,
		UserRouter: userRouter,
	}
}

// Register routes
func (r *Router) Register() {
	// Test Routes
	r.App.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Pong! 👋")
	})

	//Swagger Documentation
	r.App.Get("/swagger/*", swagger.HandlerDefault)

	// Register routes of modules
	r.UserRouter.RegisterUserRoutes()
	//r.AuthRouter.RegisterAuthRoutes()
}
