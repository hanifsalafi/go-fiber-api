package request

import (
	"go-fiber-api/app/database/entity"
	"go-fiber-api/utils/paginator"
)

type UserGeneric interface {
	ToEntity()
}

type UserQueryRequest struct {
	Username    string                `json:"username" validate:"required,lowercase"`
	Email       string                `json:"email" validate:"required,email"`
	Password    string                `json:"password" validate:"required,min=6"`
	Fullname    string                `json:"fullname" validate:"required"`
	Address     string                `json:"address" validate:"required"`
	PhoneNumber string                `json:"phoneNumber" validate:"required"`
	UserRoleID  int                   `json:"userRoleId" validate:"required,numeric"`
	Pagination  *paginator.Pagination `json:"pagination"`
}

type UserCreateRequest struct {
	Username    string `json:"username" validate:"required,lowercase"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=6"`
	Fullname    string `json:"fullname" validate:"required"`
	Address     string `json:"address" validate:"required"`
	PhoneNumber string `json:"phoneNumber" validate:"required"`
	UserRoleID  int    `json:"userRoleId" validate:"required,numeric"`
}

func (req UserCreateRequest) ToEntity() *entity.User {
	return &entity.User{
		Username:    req.Username,
		Email:       req.Email,
		Fullname:    req.Fullname,
		Address:     req.Address,
		PhoneNumber: req.PhoneNumber,
		UserRoleID:  req.UserRoleID,
		IsActive:    true,
		StatusID:    1,
	}
}

type UserUpdateRequest struct {
	ID          uint   `json:"id" validate:"required"`
	Username    string `json:"username" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=6"`
	Fullname    string `json:"fullname" validate:"required"`
	Address     string `json:"address" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	UserRoleID  int    `json:"user_role_id" validate:"required,numeric"`
}

func (req UserUpdateRequest) ToEntity() *entity.User {
	return &entity.User{
		Username:    req.Username,
		Email:       req.Email,
		Fullname:    req.Fullname,
		Address:     req.Address,
		PhoneNumber: req.PhoneNumber,
		UserRoleID:  req.UserRoleID,
	}
}
