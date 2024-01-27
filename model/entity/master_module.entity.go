package entity

import "time"

type MasterModule struct {
	ID          uint      `json:"id" gorm:"primaryKey;type:int4;autoIncrement"`
	Name        string    `json:"name" gorm:"type:varchar"`
	Description string    `json:"description" gorm:"type:varchar"`
	PathURL     string    `json:"path_url" gorm:"type:varchar"`
	StatusID    int       `json:"status_id" gorm:"type:int4"`
	IsActive    bool      `json:"is_active" gorm:"default:true"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:now()"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"default:now()"`

	// foreign key
	Status MasterStatus `json:"status" gorm:"foreignKey:StatusID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
