package seeds

import (
	"go-fiber-api/app/database/entity"
	"gorm.io/gorm"
)

type MasterStatusSeeder struct{}

var masterStatus = []entity.MasterStatus{
	{
		ID:       1,
		Name:     "Waiting",
		IsActive: true,
	},
	{
		ID:       2,
		Name:     "Active",
		IsActive: true,
	},
	{
		ID:       3,
		Name:     "Inactive",
		IsActive: true,
	},
}

func (MasterStatusSeeder) Seed(conn *gorm.DB) error {
	for _, row := range masterStatus {
		if err := conn.Create(&row).Error; err != nil {
			return err
		}
	}

	return nil
}

func (MasterStatusSeeder) Count(conn *gorm.DB) (int, error) {
	var count int64
	if err := conn.Model(&entity.MasterStatus{}).Count(&count).Error; err != nil {
		return 0, err
	}

	return int(count), nil
}
