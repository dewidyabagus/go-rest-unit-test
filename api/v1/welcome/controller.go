package welcome

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) GetWelcome(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, echo.Map{"message": "welcome my apps!"})
}
