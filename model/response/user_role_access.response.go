package response

import "time"

type UserRoleAccessResponse struct {
	MenuID          uint      `json:"menu_id"`
	UserRoleID      int       `json:"user_role_id"`
	IsViewEnabled   bool      `json:"is_view_enabled"`
	IsInsertEnabled bool      `json:"is_insert_enabled"`
	IsUpdateEnabled bool      `json:"is_update_enabled"`
	IsDeleteEnabled bool      `json:"is_delete_enabled"`
	IsAdminEnabled  bool      `json:"is_admin_enabled"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
