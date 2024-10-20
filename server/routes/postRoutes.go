package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zedann/ecoforum/server/internal/post"
	"github.com/zedann/ecoforum/server/middlewares"
)

func HandlePostRoutes(postHandler *post.PostHandler, postRouter fiber.Router) {
	postRouter.Post("/", middlewares.AuthRequire(), postHandler.CreatePost)
	postRouter.Get("/", postHandler.GetPosts)
}
