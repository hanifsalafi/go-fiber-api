package entity

import (
	"time"
)

type UserProfile struct {
	ID                 uint      `json:"id" gorm:"primaryKey;type:int4;autoIncrement"`
	Username           string    `json:"username" gorm:"type:varchar"`
	Email              string    `json:"email" gorm:"type:varchar"`
	Fullname           string    `json:"fullname" gorm:"type:varchar"`
	Address            string    `json:"address" gorm:"type:varchar"`
	PhoneNumber        string    `json:"phone_number" gorm:"type:varchar"`
	KeycloakID         string    `json:"keycloak_id" gorm:"type:varchar"`
	UserRoleID         int       `json:"user_role_id" gorm:"type:int4"`
	StatusID           int       `json:"status_id" gorm:"type:int4"`
	ProfilePicturePath string    `json:"profile_picture_path" gorm:"type:varchar"`
	IsActive           bool      `json:"is_active" gorm:"default:false"`
	CreatedAt          time.Time `json:"created_at" gorm:"default:now()"`
	UpdatedAt          time.Time `json:"updated_at" gorm:"default:now()"`

	// foreign key
	UserRole UserRole     `json:"user_role" gorm:"foreignKey:UserRoleID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Status   MasterStatus `json:"status" gorm:"foreignKey:StatusID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
