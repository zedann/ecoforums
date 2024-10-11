package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/zedann/ecoforum/server/db"
	"github.com/zedann/ecoforum/server/internal/user"
)

func main() {
	godotenv.Load()
	app := fiber.New()

	database, err := db.New()

	if err != nil {
		log.Fatal("database connection failed", err)
	}

	api := app.Group("/api/v1")

	userRepo := user.NewUserRepository(database.GetDB())
	userSvc := user.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userSvc)

	api.Post("/users", userHandler.CreateUser)
	api.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("hello")
	})

	port := os.Getenv("PORT")
	app.Listen(":" + port)
}
