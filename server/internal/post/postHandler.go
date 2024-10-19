package post

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/zedann/ecoforum/server/types"
)

type PostHandler struct {
	*PostService
}

func NewPostHandler(postSvc *PostService) *PostHandler {
	return &PostHandler{
		PostService: postSvc,
	}
}

func (h *PostHandler) CreatePost(c *fiber.Ctx) error {
	p := new(Post)

	if err := c.BodyParser(p); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":  err.Error(),
			"status": http.StatusInternalServerError,
		})
	}

	userID, err := strconv.ParseInt((c.Locals("user_id").(string)), 10, 64)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":  "can not get the userID from the middleware",
			"status": http.StatusInternalServerError,
		})
	}
	p.UserID = userID

	file, _ := c.FormFile("image")

	image := ""
	if file != nil {
		image = fmt.Sprintf("post_%s", file.Filename)
		imagePath := fmt.Sprintf("./public/imgs/post_%s", image)
		if err := c.SaveFile(file, imagePath); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error":  "Failed to save image" + err.Error(),
				"status": http.StatusInternalServerError,
			})
		}
	}

	req := &CreatePostReq{
		Title:   p.Title,
		Content: p.Content,
		Image:   image,
		UserID:  p.UserID,
	}

	// Store Post in DB
	post, err := h.PostService.CreatePost(c.Context(), req)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": http.StatusOK,
		"data":   post,
	})

}

func (h *PostHandler) GetPosts(c *fiber.Ctx) error {
	pageSize, err := strconv.Atoi(c.Params("pageSize"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":  "pageSize param should be an number",
			"status": http.StatusBadRequest,
		})
	}
	page, err := strconv.Atoi(c.Params("page"))

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":  "page param should be an number",
			"status": http.StatusBadRequest,
		})
	}
	reqConfig := types.NewReqConfig(pageSize, page)
	posts, err := h.PostService.GetPosts(c.Context(), reqConfig)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":  err.Error(),
			"status": http.StatusBadRequest,
		})

	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"data":   posts,
		"status": http.StatusOK,
	})

}
