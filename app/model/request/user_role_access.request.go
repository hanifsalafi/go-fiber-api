package request

type UserRoleAccessCreateRequest struct {
	MenuID          uint `json:"menu_id" validate:"required,numeric"`
	IsViewEnabled   bool `json:"is_view_enabled" validate:"required"`
	IsInsertEnabled bool `json:"is_insert_enabled" validate:"required"`
	IsUpdateEnabled bool `json:"is_update_enabled" validate:"required"`
	IsDeleteEnabled bool `json:"is_delete_enabled" validate:"required"`
	IsAdminEnabled  bool `json:"is_admin_enabled" validate:"required"`
}

type UserRoleAccessUpdateRequest struct {
	MenuID          uint `json:"menu_id" validate:"required,numeric"`
	UserRoleID      int  `json:"user_role_id" validate:"required,numeric"`
	IsViewEnabled   bool `json:"is_view_enabled" validate:"required"`
	IsInsertEnabled bool `json:"is_insert_enabled" validate:"required"`
	IsUpdateEnabled bool `json:"is_update_enabled" validate:"required"`
	IsDeleteEnabled bool `json:"is_delete_enabled" validate:"required"`
	IsAdminEnabled  bool `json:"is_admin_enabled" validate:"required"`
}
