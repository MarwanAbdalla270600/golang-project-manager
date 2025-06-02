package main

//connection string
// "postgres://postgres:postgres@localhost:5432/myapp?sslmode=disable"
import (
	"webapi/internal/database"
	"webapi/internal/user"

	"github.com/labstack/echo"
)

func initRoutes(app *echo.Echo, userHandler *user.UserHandler) {
	app.POST("/users", userHandler.HandleCreateUser)
}

func main() {
	database.Init()
	app := echo.New()

	//init my Singletons
	userRepo := user.NewUserRepository(database.Queries)
	userService := user.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userService)

	initRoutes(app, userHandler)

	app.Logger.Fatal(app.Start(":8080"))
}
