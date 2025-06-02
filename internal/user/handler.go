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
	var user UserInput

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadGateway, echo.Map{"error": "Invalid Request"})
	}

	

	c.JSON(http.StatusOK, UserOutput{
		Username:  user.Username,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Role:      "MEMBER",
	})
	return nil
}
