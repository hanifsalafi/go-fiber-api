package entity

import "time"

type UserRoleAccess struct {
	ID              uint      `json:"id" gorm:"primaryKey;type:int4;autoIncrement"`
	UserRoleID      int       `json:"user_role_id" gorm:"type:int4"`
	MenuID          int       `json:"menu_id" gorm:"type:int4"`
	IsViewEnabled   bool      `json:"is_view_enabled" gorm:"default:false"`
	IsInsertEnabled bool      `json:"is_insert_enabled" gorm:"default:false"`
	IsUpdateEnabled bool      `json:"is_update_enabled" gorm:"default:false"`
	IsDeleteEnabled bool      `json:"is_delete_enabled" gorm:"default:false"`
	IsAdminEnabled  bool      `json:"is_admin_enabled" gorm:"default:false"`
	IsActive        bool      `json:"is_active" gorm:"default:true"`
	CreatedAt       time.Time `json:"created_at" gorm:"default:now()"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"default:now()"`

	// foreign key
	Role UserRole   `json:"role" gorm:"foreignKey:UserRoleID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Menu MasterMenu `json:"status" gorm:"foreignKey:MenuID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
