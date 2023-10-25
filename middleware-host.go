package tootecho

import (
	"github.com/labstack/echo/v4"
)

// WithHost middleware guarantees that all incoming HTTP requests include the correct "Host" value in their headers
func WithHost(next echo.HandlerFunc) echo.HandlerFunc {

	return func(ctx echo.Context) error {
		ctx.Request().Header.Set("Host", ctx.Request().Host)
		return next(ctx)
	}
}
