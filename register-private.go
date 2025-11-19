package tootecho

import (
	"net/http"

	"github.com/benpate/derp"
	"github.com/benpate/rosetta/list"
	"github.com/benpate/toot"
	"github.com/benpate/toot/scope"
	"github.com/labstack/echo/v4"
)

// echoMethod represents an e.GET, e.POST, e.PUT, e.DELETE method that registers
// a new echo.HandlerFunc with the echo router.
type echoMethod func(string, echo.HandlerFunc, ...echo.MiddlewareFunc) *echo.Route

// commonWrapper is a function that can be used to handle both single results and paged results.
type commonWrapper[AuthToken toot.ScopesGetter, Input any, Output any] func(echo.Context, AuthToken, Input) (Output, error)

// single_result inserts a new echo.HandlerFunc into the echo router
// that returns a single result object (or an array without any paging metadata)
func single_result[AuthToken toot.ScopesGetter, Input any, Output any](api toot.API[AuthToken], fn echoMethod, path string, handler toot.APIFunc_SingleResult[AuthToken, Input, Output], requiredScope string) {

	// Do not register empty handlers
	if handler == nil {
		return
	}

	// Wrap the handler in a function that `getResult` can use
	wrapper := func(_ echo.Context, authToken AuthToken, input Input) (Output, error) {
		// Call the handler and return outputs to the caller
		return handler(authToken, input)
	}

	any_result(api, fn, path, wrapper, requiredScope)
}

// register inserts a new echo.HandlerFunc into the echo router
// that returns a paged result object.
func paged_result[AuthToken toot.ScopesGetter, Input any, Output any](api toot.API[AuthToken], fn echoMethod, path string, handler toot.APIFunc_PagedResult[AuthToken, Input, Output], requiredScope string) {

	// Do not register empty handlers
	if handler == nil {
		return
	}

	// Wrap the handler in a function that `getResult` can use
	wrapper := func(ctx echo.Context, authToken AuthToken, input Input) (Output, error) {

		// Call the actual handler
		output, pageInfo, err := handler(authToken, input)

		// Apply paging headers to the response
		pageInfo.SetHeaders(ctx.Request().Response)

		// Return outputs to the caller
		return output, err
	}

	any_result(api, fn, path, wrapper, requiredScope)
}

// any_result should not be called directly.  It is used by `single_result`
// and `paged_result` to inserts a new echo.HandlerFunc into the echo router.
// It requires a `commonWrapper` function to handle the actual request
func any_result[AuthToken toot.ScopesGetter, Input any, Output any](api toot.API[AuthToken], fn echoMethod, path string, wrapper commonWrapper[AuthToken, Input, Output], requiredScope string) {

	const location = "toot-echo.any_result"

	// If this Handler is not defined, then we don't need to register anything.
	// Calls to this route will be handled elsewhere, or will return a 404 error.
	if wrapper == nil {
		return
	}

	// Create a new echo.HandlerFunc that 1) parses inputs, 2) calls the actual handler, and
	// 3) generates a JSON response.
	tootHandler := func(ctx echo.Context) error {

		// Parse inputs from the request
		authToken, input, err := getInputs[AuthToken, Input](ctx, api, requiredScope)

		if err != nil {
			return derp.Wrap(err, location, "Error parsing inputs")
		}

		// Call the actual handler to map the Inputs to Outputs
		result, err := wrapper(ctx, authToken, input)

		if err != nil {
			return derp.Wrap(err, location, "Error executing API call")
		}

		// Set CORS header for the Mastodon API.
		ctx.Response().Header().Set("Access-Control-Allow-Origin", "*")

		// Return the API result to the caller as JSON
		if err := ctx.JSON(http.StatusOK, result); err != nil {
			return derp.Wrap(err, location, "Error writing response body")
		}

		// Woot.
		return nil
	}

	// Register the new echo.HandlerFunc with the echo Router
	fn(path, tootHandler, WithHost)
}

func getInputs[AuthToken toot.ScopesGetter, Input any](ctx echo.Context, api toot.API[AuthToken], requiredScope string) (AuthToken, Input, error) {

	const location = "toot-echo.getInputs"

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
			return authToken, input, derp.Wrap(err, location, "Request is not authorized. LOL.")
		}

		// Verify the scopes required for this API call
		if !verifyScope(authToken.Scopes(), requiredScope) {
			return authToken, input, derp.UnauthorizedError(location, "Request is not authorized.", requiredScope, authToken.Scopes())
		}
	}

	// Collect input arguments from the Request
	// TODO: HIGH: Replace this bind with custom binder:
	// https://github.com/go-playground/form
	// https://echo.labstack.com/docs/binding#custom-binding
	binder := echo.DefaultBinder{}
	if err := binder.Bind(&input, ctx); err != nil {
		return authToken, input, derp.Wrap(err, location, "Unable to read request body")
	}

	// Extra work to Bind headers, too
	if err := binder.BindHeaders(ctx, &input); err != nil {
		return authToken, input, derp.Wrap(err, location, "Error readin headers")
	}

	// Return success
	return authToken, input, nil
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
