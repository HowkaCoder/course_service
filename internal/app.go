package internal

import (
	"context"
	"course_service/internal/app/entity"
	"log"

	"github.com/redis/go-redis/v9"
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

func RedisInit() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("failed to connect to Redis")
	}
	return rdb
}
