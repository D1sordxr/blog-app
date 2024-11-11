package sqlite

import (
	"BlogWebApp/internal/storage/models/todo"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) {
	if err := db.AutoMigrate(&todo.Base{}, &todo.Author{}, &todo.Post{}, &todo.Tag{}); err != nil {
		panic("migration failed" + err.Error())
	}
}
