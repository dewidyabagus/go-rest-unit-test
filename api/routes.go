package api

import (
	echo "github.com/labstack/echo/v4"

	"go-api-v1/api/v1/welcome"
)

type Routes struct {
	Welcome *welcome.Controller
}

func RegisterRoutes(e *echo.Echo, routes *Routes) {
	v1 := e.Group("/v1")

	v1.GET("/welcome", routes.Welcome.GetWelcome)
}
