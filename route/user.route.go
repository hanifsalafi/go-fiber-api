package route

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-api/handler"
)

func UserRouteInit(r *fiber.App) {
	r.Get("/user", handler.UserGetAll)
	r.Post("/user", handler.UserCreate)
}
