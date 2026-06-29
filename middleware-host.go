package tootecho

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// WithHost middleware guarantees that all incoming HTTP requests include the correct "Host" value in their headers
func WithHost(next echo.HandlerFunc) echo.HandlerFunc {

	return func(ctx echo.Context) error {
		ctx.Request().Header.Set("Host", trueHostname(ctx.Request()))
		return next(ctx)
	}
}

// trueHostname returns the host name from the request, accounting for
// proxy headers (like X-Forwarded-Host).
func trueHostname(request *http.Request) string {

	// If this is a proxied request, then use the X-Forwarded-Host header
	// instead of the Host header
	if trueHost := request.Header.Get("X-Forwarded-Host"); trueHost != "" {
		return trueHost
	}

	// Fallback to the Host header if X-Forwarded-Host is not present
	return request.Host
}
