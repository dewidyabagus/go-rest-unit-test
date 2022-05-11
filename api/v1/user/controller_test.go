package user

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	echo "github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	payload = `{"email":"widya@example.com","name":"Widya Ade Bagus"}`
	email   = "widya@example.com"
)

func TestNewUser(t *testing.T) {
	e := echo.New()

	handler := NewController()

	reqError := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{""}`))
	reqError.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recError := httptest.NewRecorder()
	ctxError := e.NewContext(reqError, recError)
	assert.NoError(t, handler.NewUser(ctxError))
	assert.Equal(t, http.StatusBadRequest, recError.Code)

	for i := 1; i <= 2; i++ {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(payload))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		assert.NoError(t, handler.NewUser(ctx))

		switch i {
		case 1:
			assert.Equal(t, http.StatusCreated, rec.Code)

		case 2:
			assert.Equal(t, http.StatusConflict, rec.Code)
		}
	}

}

func TestGetWithEmail(t *testing.T) {
	e := echo.New()

	handler := NewController()

	reqError := httptest.NewRequest(http.MethodGet, "/", nil)
	recError := httptest.NewRecorder()
	ctxError := e.NewContext(reqError, recError)
	ctxError.SetPath("/:email")
	ctxError.SetParamNames("email")
	ctxError.SetParamValues("testing@email.com")

	assert.NoError(t, handler.GetWithEmail(ctxError))
	assert.Equal(t, http.StatusNotFound, recError.Code)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/:email")
	ctx.SetParamNames("email")
	ctx.SetParamValues(email)

	assert.NoError(t, handler.GetWithEmail(ctx))
	assert.Equal(t, http.StatusOK, rec.Code)

	user := new(User)
	json.Unmarshal(rec.Body.Bytes(), user)
	assert.Equal(t, "Widya Ade Bagus", user.Name)
}

func TestGetAll(t *testing.T) {
	e := echo.New()

	handler := NewController()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	if assert.NoError(t, handler.GetAll(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		user := new([]User)
		json.Unmarshal(rec.Body.Bytes(), user)
		assert.Equal(t, 1, len(*user))
	}
}

func TestUpdateWithEmail(t *testing.T) {
	e := echo.New()

	handler := NewController()

	testCases := []map[string]string{
		{
			"email":   "email@notfound.id",
			"payload": `{"invalid"}`,
			"code":    "400",
		},
		{
			"email":   "email@notfound.id",
			"payload": `{"email":"email@notfound.id","name":"name not found"}`,
			"code":    "404",
		},
		{
			"email":   email,
			"payload": payload,
			"code":    "200",
		},
	}
	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(testCase["payload"]))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		ctx := e.NewContext(req, rec)
		ctx.SetPath("/:email")
		ctx.SetParamNames("email")
		ctx.SetParamValues(testCase["email"])

		if assert.NoError(t, handler.UpdateWithEmail(ctx)) {
			assert.Equal(t, testCase["code"], strconv.Itoa(rec.Code))
		}
	}
}

func TestDeleteWithEmail(t *testing.T) {
	e := echo.New()

	handler := NewController()

	testCases := []map[string]string{
		{
			"email": "email@notfound.id",
			"code":  "404",
		},
		{
			"email": email,
			"code":  "200",
		},
	}

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		rec := httptest.NewRecorder()

		ctx := e.NewContext(req, rec)
		ctx.SetPath("/:email")
		ctx.SetParamNames("email")
		ctx.SetParamValues(testCase["email"])

		if assert.NoError(t, handler.DeleteWithEmail(ctx)) {
			assert.Equal(t, testCase["code"], strconv.Itoa(rec.Code))
		}
	}
}
