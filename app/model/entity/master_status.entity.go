package entity

type MasterStatus struct {
	ID       uint   `gorm:"primaryKey;type:int4;autoIncrement"`
	Name     string `gorm:"type:varchar"`
	IsActive bool   `gorm:"default:true"`
}
