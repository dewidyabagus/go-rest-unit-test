package welcome

import (
	"net/http"
	"net/http/httptest"
	"testing"

	echo "github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetWelcome(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	handler := NewController()
	assert.NotNil(t, handler)

	if assert.NoError(t, handler.GetWelcome(ctx)) {
		msgSuccess := `{"message":"welcome my apps!"}`
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), msgSuccess)
	}
}
