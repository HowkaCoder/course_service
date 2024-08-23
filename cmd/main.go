package main

import (
	"course_service/internal"
	"course_service/internal/app/handler"
	"course_service/internal/app/repository"
	"course_service/internal/app/usecase"

	"github.com/redis/go-redis/v9"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	RDB *redis.Client
)

func main() {

	DB = internal.DatabaseInit()

	RDB = internal.RedisInit()

	actRepository := repository.NewActRepository(DB, RDB)
	actUsecase := usecase.NewActUsecase(actRepository)
	actHandler := handler.NewActHandler(actUsecase)

	app := fiber.New()

	app.Get("/api/acts/:id", actHandler.GetActsByAnswerID)
	app.Post("/api/acts", actHandler.CreateAct)
	app.Patch("/api/acts/:id", actHandler.UpdateAct)
	app.Delete("/api/acts", actHandler.DeleteAct)

	app.Listen(":8080")

}
