package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/zedann/ecoforum/server/db"
	"github.com/zedann/ecoforum/server/internal/post"
	"github.com/zedann/ecoforum/server/internal/user"
	"github.com/zedann/ecoforum/server/routes"
)

func main() {
	godotenv.Load()
	app := fiber.New()

	database, err := db.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",        // Allow requests from your frontend
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE",     // Allowed methods
		AllowHeaders:     "Origin, Content-Type, Accept", // Allowed headers
		AllowCredentials: true,                           // Allow credentials (cookies, authorization headers, etc.)
	}))

	if err != nil {
		log.Fatal("database connection failed", err)
	}
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Static("/images", "./public/imgs")

	api := app.Group("/api/v1")

	// User Entity
	userRepo := user.NewUserRepository(database.GetDB())
	userSvc := user.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userSvc)
	userRouter := api.Group("/users")
	routes.HandleUserRoutes(userHandler, userRouter)
	// Post Entity

	postRepo := post.NewPostRepository(database.GetDB())
	postSvc := post.NewPostService(postRepo)
	postHandler := post.NewPostHandler(postSvc)
	postRouter := api.Group("/posts")
	routes.HandlePostRoutes(postHandler, postRouter)

	port := os.Getenv("PORT")
	app.Listen(":" + port)
}
