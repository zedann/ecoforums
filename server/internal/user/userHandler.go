package user

import (
	"net/http"
	"time"

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
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status": http.StatusInternalServerError,
			"error":  err.Error(),
		})
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

func (h *UserHandler) Login(c *fiber.Ctx) error {
	u := new(User)

	if err := c.BodyParser(u); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status": http.StatusInternalServerError,
			"error":  err.Error(),
		})
	}

	loginReq := &LoginUserReq{
		Email:    u.Email,
		Password: u.Password,
	}
	res, err := h.UserService.Login(c.Context(), loginReq)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status": http.StatusInternalServerError,
			"error":  err.Error(),
		})
	}

	// set cookie

	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    res.accessToken,
		Expires:  time.Now().Add(time.Hour * 24 * 10),
		HTTPOnly: true,
	})

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": http.StatusOK,
		"data":   res,
	})
}

func (h *UserHandler) Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(time.Second * 5),
		HTTPOnly: true,
	})
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": http.StatusOK,
		"data":   "Logged Out Successfuly",
	})
}

func (h *UserHandler) Test(c *fiber.Ctx) error {
	jwtCookie := c.Cookies("jwt")
	if jwtCookie == "" {
		return c.SendString("Not Logged in")
	}
	return c.SendString("Logged in")
}
