package controller

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-api/app/module/user/request"
	"go-fiber-api/app/module/user/service"
	"go-fiber-api/utils/paginator"
	"strconv"

	utilRes "go-fiber-api/utils/response"
	utilVal "go-fiber-api/utils/validator"
)

type userController struct {
	userService service.UserService
}

type UserController interface {
	All(c *fiber.Ctx) error
	Show(c *fiber.Ctx) error
	Save(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

func NewUserController(userService service.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

// All get all articles
// @Summary      Get all articles
// @Description  API for getting all articles
// @Tags         Task
// @Security     Bearer
// @Success      200  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      422  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /articles [get]
func (_i *userController) All(c *fiber.Ctx) error {
	paginate, err := paginator.Paginate(c)
	if err != nil {
		return err
	}

	var req request.UserQueryRequest
	req.Pagination = paginate

	articles, paging, err := _i.userService.All(req)
	if err != nil {
		return err
	}

	return utilRes.Resp(c, utilRes.Response{
		Messages: utilRes.Messages{"Article list successfully retrieved"},
		Data:     articles,
		Meta:     paging,
	})
}

// Show get one article
// @Summary      Get one article
// @Description  API for getting one article
// @Tags         Task
// @Security     Bearer
// @Param        id path int true "Article ID"
// @Success      200  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      422  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /articles/:id [get]
func (_i *userController) Show(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 0)
	if err != nil {
		return err
	}

	articles, err := _i.userService.Show(uint(id))
	if err != nil {
		return err
	}

	return utilRes.Resp(c, utilRes.Response{
		Messages: utilRes.Messages{"Article successfully retrieved"},
		Data:     articles,
	})
}

// Save create article
// @Summary      Create article
// @Description  API for create article
// @Tags         Task
// @Security     Bearer
// @Body 	     request.ArticleRequest
// @Success      200  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      422  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /articles [post]
func (_i *userController) Save(c *fiber.Ctx) error {
	req := new(request.UserCreateRequest)
	if err := utilVal.ParseAndValidate(c, req); err != nil {
		return err
	}

	err := _i.userService.Save(*req)
	if err != nil {
		return err
	}

	return utilRes.Resp(c, utilRes.Response{
		Messages: utilRes.Messages{"Article successfully created"},
	})
}

// Update update article
// @Summary      update article
// @Description  API for update article
// @Tags         Task
// @Security     Bearer
// @Body 	     request.ArticleRequest
// @Param        id path int true "Article ID"
// @Success      200  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      422  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /articles/:id [put]
func (_i *userController) Update(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 0)
	if err != nil {
		return err
	}

	req := new(request.UserUpdateRequest)
	if err := utilVal.ParseAndValidate(c, req); err != nil {
		return err
	}

	err = _i.userService.Update(uint(id), *req)
	if err != nil {
		return err
	}

	return utilRes.Resp(c, utilRes.Response{
		Messages: utilRes.Messages{"Article successfully updated"},
	})
}

// Delete delete article
// @Summary      delete article
// @Description  API for delete article
// @Tags         Task
// @Security     Bearer
// @Param        id path int true "Article ID"
// @Success      200  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      422  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /articles/:id [delete]
func (_i *userController) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 0)
	if err != nil {
		return err
	}

	err = _i.userService.Delete(uint(id))
	if err != nil {
		return err
	}

	return utilRes.Resp(c, utilRes.Response{
		Messages: utilRes.Messages{"Article successfully deleted"},
	})
}
