package mapper

import (
	"go-fiber-api/app/database/entity"
	res "go-fiber-api/app/module/user/response"
)

func UserResponseMapper(userReq *entity.User) (userRes *res.UserResponse) {
	if userReq != nil {
		userRes = &res.UserResponse{
			ID:          userReq.ID,
			Username:    userReq.Username,
			Email:       userReq.Email,
			Fullname:    userReq.Fullname,
			Address:     userReq.Address,
			PhoneNumber: userReq.PhoneNumber,
			UserRoleID:  userReq.UserRoleID,
			StatusID:    userReq.StatusID,
			CreatedAt:   userReq.CreatedAt,
			UpdatedAt:   userReq.UpdatedAt,
		}
	}
	return userRes
}
