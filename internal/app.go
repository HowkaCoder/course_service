package internal

import (
	"course_service/internal/app/entity"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func DatabaseInit() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("database/database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&entity.Act{})

	return db

}
