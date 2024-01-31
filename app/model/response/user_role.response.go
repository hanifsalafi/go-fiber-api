package response

import "time"

type UserRoleResponse struct {
	ID             uint                     `json:"id"`
	Name           string                   `json:"name"`
	Description    string                   `json:"description"`
	PathURL        string                   `json:"path_url"`
	StatusID       int                      `json:"status_id"`
	UserRoleAccess []UserRoleAccessResponse `json:"user_role_access"`
	CreatedAt      time.Time                `json:"created_at"`
	UpdatedAt      time.Time                `json:"updated_at"`
}
