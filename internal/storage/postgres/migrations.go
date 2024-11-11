package postgres

import (
	"BlogWebApp/internal/storage/models"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) {
	if err := db.AutoMigrate(&models.Post{}); err != nil {
		panic("migration failed" + err.Error())
	}
}
