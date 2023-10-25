package tootecho

import (
	"net/http"

	"github.com/benpate/derp"
	"github.com/benpate/rosetta/list"
	"github.com/benpate/toot"
	"github.com/benpate/toot/scope"
	"github.com/labstack/echo/v4"
)

// register inserts a new echo.HandlerFunc into the echo router.
func register[AuthToken toot.ScopesGetter, Input any, Output any](api toot.API[AuthToken], fn echoMethod, path string, handler toot.APIFunc[AuthToken, Input, Output], requiredScope string) {

	const location = "toot-echo.register"

	// If this Handler is not defined, then skip it.
	if handler == nil {
		return
	}

	// Register a new echo.HandlerFunc in the echo Router
	fn(path, func(ctx echo.Context) error {

		var input Input
		var authToken AuthToken
		var err error

		// If the request is not public (at least one scope is required)
		// then try to authorize the request.
		// If no scopes are required, then an empty AuthToken
		// will be passed to the handler.
		if requiredScope != scope.Public {

			authToken, err = api.Authorize(ctx.Request())

			if err != nil {
				return derp.Wrap(err, location, "Request is not authorized. LOL.")
			}

			// Verify the scopes required for this API call
			if !verifyScope(authToken.Scopes(), requiredScope) {
				return derp.NewUnauthorizedError(location, "Request is not authorized.", requiredScope, authToken.Scopes())
			}
		}

		// Collect input arguments from the Request
		// TODO: HIGH: Replace this bind with custom binder:
		// https://github.com/go-playground/form
		// https://echo.labstack.com/docs/binding#custom-binding
		binder := echo.DefaultBinder{}
		if err := binder.Bind(&input, ctx); err != nil {
			return derp.Wrap(err, location, "Error reading request body")
		}

		// Extra work to Bind headers, too
		if err := binder.BindHeaders(ctx, &input); err != nil {
			return derp.Wrap(err, location, "Error readin headers")
		}

		// Execute the API handler
		result, err := handler(authToken, input)

		if err != nil {
			return derp.Wrap(err, location, "Error executing API call")
		}

		// Return the API result to the caller as JSON
		if err := ctx.JSON(http.StatusOK, result); err != nil {
			return derp.Wrap(err, location, "Error writing response body")
		}

		// Woot.
		return nil
	}, WithHost)

}

// verifyScope confirms that the required scope exists in the
// `present` slice.
func verifyScope(present []string, requiredScope string) bool {

	// Always allow public requests
	if requiredScope == scope.Public {
		return true
	}

	// Since we're already authenticated, "private" requests
	// with no additional scope requirements are also allowed
	if requiredScope == scope.Private {
		return true
	}

	// If the required scope contains a colon, see if the user has just the "prefix" scope
	if prefix, suffix := list.Split(requiredScope, ':'); suffix != "" {
		for _, scope := range present {
			if scope == prefix {
				return true
			}
		}
	}

	// Otherwise, search for the full scope in the `present` list
	for _, scope := range present {
		if scope == requiredScope {
			return true
		}
	}

	// No scope was found in the `present` list. This request will be denied.
	return false
}
