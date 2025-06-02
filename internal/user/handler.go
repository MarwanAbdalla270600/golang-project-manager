package user

import (
	"net/http"
	"webapi/internal/db"

	"github.com/labstack/echo"
)

type UserHandler struct {
	UserService *UserService
}

func NewUserHandler(UserService *UserService) *UserHandler {
	return &UserHandler{UserService: UserService}
}

func (h *UserHandler) HandleCreateUser(c echo.Context) error {
	var user db.User

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadGateway, echo.Map{"error": "Invalid Request"})
	}

	c.JSON(http.StatusOK, "created user")
	return nil
}
