package response

import "time"

type UserProfileResponse struct {
	ID          uint      `json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Fullname    string    `json:"fullname"`
	Address     string    `json:"address"`
	PhoneNumber string    `json:"phone_number"`
	UserRoleID  int       `json:"user_role_id"`
	StatusID    int       `json:"status_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
