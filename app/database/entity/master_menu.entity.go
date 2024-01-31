package entity

import "time"

type MasterMenu struct {
	ID           uint      `json:"id" gorm:"primaryKey;type:int4;autoIncrement"`
	Name         string    `json:"name" gorm:"type:varchar"`
	Description  string    `json:"description" gorm:"type:varchar"`
	ModuleID     int       `json:"module_id" gorm:"type:int4"`
	ParentMenuID int       `json:"parent_menu_id" gorm:"type:int4"`
	Icon         string    `json:"icon" gorm:"type:varchar"`
	Position     int       `json:"position" gorm:"type:int4"`
	StatusID     int       `json:"status_id" gorm:"type:int4"`
	IsActive     bool      `json:"is_active" gorm:"default:true"`
	CreatedAt    time.Time `json:"created_at" gorm:"default:now()"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"default:now()"`

	// foreign key
	Module MasterModule `json:"module" gorm:"foreignKey:ModuleID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Status MasterStatus `json:"status" gorm:"foreignKey:StatusID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
