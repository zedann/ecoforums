package user

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	*UserService
}

func NewUserHandler(userSvc *UserService) *UserHandler {
	return &UserHandler{
		UserService: userSvc,
	}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	u := new(User)

	if err := c.BodyParser(u); err != nil {
		return err
	}

	if u.Username == "" || u.Email == "" || u.Password == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status": http.StatusBadRequest,
			"error":  "Please Fill All Required Fields",
		})

	}

	req := &CreateUserReq{
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
	}
	user, err := h.UserService.CreateUser(c.Context(), req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": http.StatusOK,
		"data":   user,
	})

}
