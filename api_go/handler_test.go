package demo_test

import (
	"demo"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHandlerSayHi(t *testing.T) {
	// Arrange
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/hello")

	// Mock Database
	message := demo.Message{Data: "Hello from mock database"}

	// Assertions
	var messageJSON = `{"message":"Hello from mock database"}`
	f := demo.SayHi(message)
	if assert.NoError(t, f(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, messageJSON, strings.TrimSpace(rec.Body.String()))
	}
}
