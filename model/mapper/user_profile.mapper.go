package mapper

import (
	e "go-fiber-api/model/entity"
	res "go-fiber-api/model/response"
)

func UserProfileResponseMapper(user *e.UserProfile) res.UserProfileResponse {
	userProfile := res.UserProfileResponse{
		ID:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		Fullname:    user.Fullname,
		Address:     user.Address,
		PhoneNumber: user.PhoneNumber,
		UserRoleID:  user.UserRoleID,
		StatusID:    user.StatusID,
	}
	return userProfile
}
