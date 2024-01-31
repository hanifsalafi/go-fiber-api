package controller

import "go-fiber-api/app/module/user/service"

type Controller struct {
	User UserController
}

func NewController(articleService service.UserService) *Controller {
	return &Controller{
		User: NewUserController(articleService),
	}
}
