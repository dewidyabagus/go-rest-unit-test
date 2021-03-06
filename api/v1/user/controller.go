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

var (
	DB = map[string]*User{}

	successResponse      = echo.Map{"message": "success"}
	errInvalidBodyFormat = echo.Map{"message": "invalid body format"}
	errUserNotFound      = echo.Map{"message": "user record not found"}
)

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) NewUser(ctx echo.Context) error {
	user := new(User)

	if err := ctx.Bind(user); err != nil {
		return ctx.JSON(http.StatusBadRequest, errInvalidBodyFormat)
	}

	if _, found := DB[user.Email]; found {
		return ctx.JSON(http.StatusConflict, echo.Map{"message": "data already exists"})
	}

	DB[user.Email] = user

	return ctx.JSON(http.StatusCreated, successResponse)
}

func (c *Controller) GetAll(ctx echo.Context) error {
	users := []User{}

	for _, user := range DB {
		users = append(users, *user)
	}

	return ctx.JSON(http.StatusOK, users)
}

func (c *Controller) GetWithEmail(ctx echo.Context) error {
	email := ctx.Param("email")

	user, found := DB[email]
	if !found {
		return ctx.JSON(http.StatusNotFound, errUserNotFound)
	}

	return ctx.JSON(http.StatusOK, user)
}

func (c *Controller) UpdateWithEmail(ctx echo.Context) error {
	user := new(User)
	email := ctx.Param("email")

	if err := ctx.Bind(user); err != nil {
		return ctx.JSON(http.StatusBadRequest, errInvalidBodyFormat)
	}

	if _, found := DB[email]; !found {
		return ctx.JSON(http.StatusNotFound, errUserNotFound)
	}

	delete(DB, email)
	DB[user.Email] = user

	return ctx.JSON(http.StatusOK, successResponse)
}

func (c *Controller) DeleteWithEmail(ctx echo.Context) error {
	email := ctx.Param("email")

	if _, found := DB[email]; !found {
		return ctx.JSON(http.StatusNotFound, errUserNotFound)
	}

	delete(DB, email)

	return ctx.JSON(http.StatusOK, successResponse)
}
