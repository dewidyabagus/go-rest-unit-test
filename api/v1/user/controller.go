package user

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type (
	Controller struct {
	}

	User struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	}
)

var DB = map[string]*User{}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) NewUser(ctx echo.Context) error {
	user := new(User)

	if err := ctx.Bind(user); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": "invalid body format"})
	}

	if _, found := DB[user.Email]; found {
		return ctx.JSON(http.StatusConflict, echo.Map{"message": "data already exists"})
	}

	DB[user.Email] = user

	return ctx.JSON(http.StatusCreated, echo.Map{"message": "success"})
}

func (c *Controller) GetWithEmail(ctx echo.Context) error {
	email := ctx.Param("email")

	user, found := DB[email]
	if !found {
		return ctx.JSON(http.StatusNotFound, echo.Map{"message": "user record not found"})
	}

	return ctx.JSON(http.StatusOK, user)
}
