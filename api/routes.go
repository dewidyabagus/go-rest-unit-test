package api

import (
	echo "github.com/labstack/echo/v4"

	"go-api-v1/api/v1/user"
	"go-api-v1/api/v1/welcome"
)

type Routes struct {
	Welcome *welcome.Controller
	User    *user.Controller
}

func RegisterRoutes(e *echo.Echo, routes *Routes) {
	v1 := e.Group("/v1")

	v1.GET("/welcome", routes.Welcome.GetWelcome)
	v1.POST("/users", routes.User.NewUser)
	v1.GET("/users", routes.User.GetAll)
	v1.GET("/users/:email", routes.User.GetWithEmail)
	v1.PUT("/users/:email", routes.User.UpdateWithEmail)
	v1.DELETE("/users/:email", routes.User.DeleteWithEmail)
}
