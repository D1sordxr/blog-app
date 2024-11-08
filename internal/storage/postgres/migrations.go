package postgres

import (
	"BlogWebApp/internal/storage/models"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) {
	if err := db.AutoMigrate(&models.Base{}, &models.Author{}, &models.Post{}, &models.Tag{}); err != nil {
		panic("migration failed" + err.Error())
	}
}
