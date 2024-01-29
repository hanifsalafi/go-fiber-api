package response

import "time"

type MasterMenuResponse struct {
	ID           uint      `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	ModuleID     int       `json:"module_id"`
	ParentMenuID int       `json:"parent_menu_id"`
	Icon         string    `json:"icon"`
	Position     int       `json:"position"`
	StatusID     int       `json:"status_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
