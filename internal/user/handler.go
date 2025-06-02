package user

import (
	"net/http"

	"github.com/labstack/echo"
)

type UserHandler struct {
	UserService *UserService
}

func NewUserHandler(UserService *UserService) *UserHandler {
	return &UserHandler{UserService: UserService}
}

func (h *UserHandler) HandleCreateUser(c echo.Context) error {
	var userInput *UserInput

	if err := c.Bind(userInput); err != nil {
		return c.JSON(http.StatusBadGateway, echo.Map{"error": "Invalid Request"})
	}

	userOutput, err := h.UserService.CreateUser(userInput)
	if err != nil {
		return c.JSON(http.StatusBadGateway, echo.Map{"error": "Invalid Request"})
	}

	c.JSON(http.StatusOK, userOutput)
	return nil
}
