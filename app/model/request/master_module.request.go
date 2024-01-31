package request

type MasterModuleCreateRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	PathURL     string `json:"path_url" validate:"required"`
}

type MasterModuleUpdateRequest struct {
	ID          uint   `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	PathURL     string `json:"path_url" validate:"required"`
	StatusID    int    `json:"status_id" validate:"required,numeric"`
}
