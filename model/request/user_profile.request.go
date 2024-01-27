package request

type UserProfileCreateRequest struct {
	Username    string `json:"username" validate:"required,lowercase"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=6"`
	Fullname    string `json:"fullname" validate:"required"`
	Address     string `json:"address" validate:"required"`
	PhoneNumber string `json:"phoneNumber" validate:"required"`
	UserRoleID  int    `json:"userRoleId" validate:"required,numeric"`
}

type UserProfileUpdateRequest struct {
	ID       uint   `json:"id" validate:"required"`
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Fullname string `json:"fullname" validate:"required"`
	Address  string `json:"address" validate:"required"`

	PhoneNumber string `json:"phone_number" validate:"required"`
	UserRoleID  int    `json:"user_role_id" validate:"required,numeric"`
	StatusID    int    `json:"status_id" validate:"required,numeric"`
}
