package request

type UserRoleCreateRequest struct {
	Name           string                        `json:"name" validate:"required"`
	Description    string                        `json:"description" validate:"required"`
	Code           string                        `json:"code" validate:"required"`
	UserRoleAccess []UserRoleAccessCreateRequest `json:"user_role_access" validate:"required"`
}

type UserRoleUpdateRequest struct {
	ID             uint                          `json:"id" validate:"required"`
	Name           string                        `json:"name" validate:"required"`
	Description    string                        `json:"description" validate:"required"`
	Code           string                        `json:"code" validate:"required"`
	StatusID       int                           `json:"status_id" validate:"required,numeric"`
	UserRoleAccess []UserRoleAccessUpdateRequest `json:"user_role_access" validate:"required"`
}
