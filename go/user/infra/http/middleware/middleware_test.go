package middleware_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"sample/user/infra/http/middleware"
	"slices"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestUse(t *testing.T) {
	e := echo.New()
	middleware.Use(e)

	// Disabled to write logs
	e.Pre(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Logger().SetOutput(io.Discard)
			return next(c)
		}
	})

	req := httptest.NewRequest(http.MethodOptions, "/", nil)
	req.Header.Set(echo.HeaderOrigin, "http://example.com")
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	t.Run("should AccessControlAllowOrigin successflly", func(t *testing.T) {
		expected := "*"
		actual := rec.Header().Get(echo.HeaderAccessControlAllowOrigin)

		if actual != expected {
			t.Errorf("AccessControlAllowOrigin missmatch: expected %q, but got %q", expected, actual)
		}
	})

	methods := []string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodDelete,
	}

	for _, m := range methods {
		t.Run(fmt.Sprintf("should include AllowMethod %q successfully", m), func(t *testing.T) {
			str := rec.Header().Get(echo.HeaderAccessControlAllowMethods)
			arr := strings.Split(str, ",")

			if !slices.Contains(arr, m) {
				t.Errorf("method not include: %q, but methods is %q", m, str)
			}
		})
	}

	t.Run("should include Origin header successfully", func(t *testing.T) {
		str := rec.Header().Get(echo.HeaderAccessControlAllowHeaders)
		arr := strings.Split(str, ",")

		if !slices.Contains(arr, echo.HeaderOrigin) {
			t.Errorf("Origin header not include: %q, but Origin header is %q", echo.HeaderOrigin, str)
		}
	})
}
