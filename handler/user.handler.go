package handler

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-api/app/database"
	e "go-fiber-api/model/entity"
	"go-fiber-api/model/mapper"
	req "go-fiber-api/model/request"
	res "go-fiber-api/model/response"
	"go-fiber-api/utils"
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

	// Validate Request
	if err := validator.ValidateRequest(user); err != nil {
		returnData := res.ReturnData{
			Success: false,
			Data:    err.Error(),
			Message: "Validation Error",
		}
		return ctx.Status(400).JSON(returnData)
	}

	// Save Model
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

	userResponse := result.Statement.Model.(e.UserProfile)

	// Mapping Response
	resultResponse := mapper.UserProfileResponseMapper(userResponse)

	returnData := res.ReturnData{
		Success: true,
		Data:    resultResponse,
		Message: "Success",
	}

	return ctx.JSON(returnData)
}

func UserUpdate(ctx *fiber.Ctx) error {
	userRequest := new(req.UserProfileUpdateRequest)
	if err := ctx.BodyParser(userRequest); err != nil {
		return err
	}

	// Validate Request
	if err := validator.ValidateRequest(userRequest); err != nil {
		returnData := res.ReturnData{
			Success: false,
			Data:    err.Error(),
			Message: "Validation Error",
		}
		return ctx.Status(400).JSON(returnData)
	}

	var user e.UserProfile
	// Find User
	err := database.DB.First(&user, "id = ?", user.ID).Error
	if err != nil {
		returnData := res.ReturnData{
			Success: false,
			Data:    "User Not Found",
			Message: "User Not Found",
		}

		return ctx.Status(404).JSON(returnData)
	}

	// Update User
	user.Username = userRequest.Username
	user.Email = userRequest.Email
	user.Fullname = userRequest.Fullname
	user.Address = userRequest.Address
	user.PhoneNumber = userRequest.PhoneNumber
	user.UserRoleID = userRequest.UserRoleID

	errUpdate := database.DB.Save(&user).Error
	if errUpdate != nil {
		returnData := res.ReturnData{
			Success: false,
			Data:    "Save User Failed",
			Message: "Internal Server Error",
		}

		return ctx.Status(500).JSON(returnData)
	}

	// Mapping Response
	resultResponse := mapper.UserProfileResponseMapper(user)

	returnData := res.ReturnData{
		Success: true,
		Data:    resultResponse,
		Message: "Success",
	}

	return ctx.JSON(returnData)
}
