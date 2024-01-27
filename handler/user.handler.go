package handler

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-api/database"
	e "go-fiber-api/model/entity"
	"go-fiber-api/model/mapper"
	req "go-fiber-api/model/request"
	res "go-fiber-api/model/response"
	"go-fiber-api/model/validator"
	"log"
)

func UserGetAll(ctx *fiber.Ctx) error {
	var users []e.UserProfile

	result := database.DB.Preload("UserRole").Preload("Status").Debug().Find(&users)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(users)
}

func UserCreate(ctx *fiber.Ctx) error {
	user := new(req.UserProfileCreateRequest)
	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	// Validate request
	if err := validator.ValidateUserRequest(user); err != nil {
		returnData := res.ReturnData{
			Success: false,
			Data:    err.Error(),
			Message: "Validation Error",
		}
		return ctx.Status(400).JSON(returnData)
	}

	newUser := e.UserProfile{
		Username:    user.Username,
		Email:       user.Email,
		Fullname:    user.Fullname,
		Address:     user.Address,
		PhoneNumber: user.PhoneNumber,
		UserRoleID:  user.UserRoleID,
		IsActive:    true,
		StatusID:    1,
	}

	result := database.DB.Create(&newUser)
	if result.Error != nil {
		returnData := res.ReturnData{
			Success: true,
			Data:    result.Error,
			Message: "Success",
		}

		return ctx.Status(500).JSON(returnData)
	}

	userResponse := result.Statement.Model.(*e.UserProfile)

	//
	resultResponse := mapper.UserProfileResponseMapper(userResponse)

	returnData := res.ReturnData{
		Success: true,
		Data:    resultResponse,
		Message: "Success",
	}

	return ctx.JSON(returnData)
}
