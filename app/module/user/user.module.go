package user

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-api/app/module/user/controller"
	"go-fiber-api/app/module/user/repository"
	"go-fiber-api/app/module/user/service"
	"go.uber.org/fx"
)

// struct of ArticleRouter
type UserRouter struct {
	App        fiber.Router
	Controller *controller.Controller
}

// register bulky of article module
var NewUserModule = fx.Options(
	// register repository of article module
	fx.Provide(repository.NewUserRepository),

	// register service of article module
	fx.Provide(service.NewArticleService),

	// register controller of article module
	fx.Provide(controller.NewController),

	// register router of article module
	fx.Provide(NewUserRouter),
)

// init ArticleRouter
func NewUserRouter(fiber *fiber.App, controller *controller.Controller) *UserRouter {
	return &UserRouter{
		App:        fiber,
		Controller: controller,
	}
}

// register routes of article module
func (_i *UserRouter) RegisterUserRoutes() {
	// define controllers
	userController := _i.Controller.User

	// define routes
	_i.App.Route("/user", func(router fiber.Router) {
		router.Get("/", userController.All)
		router.Get("/:id", userController.Show)
		router.Post("/", userController.Save)
		router.Put("/:id", userController.Update)
		router.Delete("/:id", userController.Delete)
	})
}
