package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zedann/ecoforum/server/internal/user"
)

func HandleUserRoutes(userHandler *user.UserHandler, userRouter fiber.Router) {
	userRouter.Post("/signup", userHandler.CreateUser)
	userRouter.Post("/login", userHandler.Login)
	userRouter.Get("/logout", userHandler.Logout)
	userRouter.Get("/test", userHandler.Test)
}
