package request

type MasterMenuCreateRequest struct {
	Name         string `json:"name" validate:"required"`
	Description  string `json:"description" validate:"required"`
	ModuleID     int    `json:"module_id" validate:"required,numeric"`
	ParentMenuID int    `json:"parent_menu_id" validate:"required,numeric"`
	Icon         string `json:"icon" validate:"required"`
	Position     int    `json:"position" validate:"required,numeric"`
}

type MasterMenuUpdateRequest struct {
	ID           uint   `json:"id" validate:"required"`
	Name         string `json:"name" validate:"required"`
	Description  string `json:"description" validate:"required"`
	ModuleID     int    `json:"module_id" validate:"required,numeric"`
	ParentMenuID int    `json:"parent_menu_id" validate:"required,numeric"`
	Icon         string `json:"icon" validate:"required"`
	Position     int    `json:"position" validate:"required,numeric"`
	StatusID     int    `json:"status_id" validate:"required,numeric"`
}
