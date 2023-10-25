package tootecho

import (
	"fmt"
	"net/http"

	"github.com/benpate/derp"
	"github.com/benpate/rosetta/list"
	"github.com/benpate/toot"
	"github.com/benpate/toot/route"
	"github.com/benpate/toot/scope"
	"github.com/labstack/echo/v4"
)

type echoMethod func(string, echo.HandlerFunc, ...echo.MiddlewareFunc) *echo.Route

func Register[AuthToken toot.ScopesGetter](e *echo.Echo, api toot.API[AuthToken], middleware ...echo.MiddlewareFunc) {

	fmt.Println("REGISTER: BEGIN")

	// https://docs.joinmastodon.org/methods/accounts/
	register(api, e.POST, route.PostAccount, api.PostAccount, scope.PostAccount)
	register(api, e.GET, route.GetAccount_VerifyCredentials, api.GetAccount_VerifyCredentials, scope.GetAccount_VerifyCredentials)
	register(api, e.PATCH, route.PatchAccount_UpdateCredentials, api.PatchAccount_UpdateCredentials, scope.PatchAccount_UpdateCredentials)
	register(api, e.GET, route.GetAccount, api.GetAccount, scope.GetAccount)
	register(api, e.GET, route.GetAccount_Statuses, api.GetAccount_Statuses, scope.GetAccount_Statuses)
	register(api, e.GET, route.GetAccount_Followers, api.GetAccount_Followers, scope.GetAccount_Followers)
	register(api, e.GET, route.GetAccount_Following, api.GetAccount_Following, scope.GetAccount_Following)
	register(api, e.GET, route.GetAccount_FeaturedTags, api.GetAccount_FeaturedTags, scope.GetAccount_FeaturedTags)
	register(api, e.POST, route.PostAccount, api.PostAccount_Follow, scope.PostAccont_Follow)
	register(api, e.POST, route.PostAccount_Unfollow, api.PostAccount_Unfollow, scope.PostAccount_Unfollow)
	register(api, e.POST, route.PostAccount_Block, api.PostAccount_Block, scope.PostAccount_Block)
	register(api, e.POST, route.PostAccount_Unblock, api.PostAccount_Unblock, scope.PostAccount_Unblock)
	register(api, e.POST, route.PostAccount_Mute, api.PostAccount_Mute, scope.PostAccount_Mute)
	register(api, e.POST, route.PostAccount_Unmute, api.PostAccount_Unmute, scope.PostAccount_Unmute)
	register(api, e.POST, route.PostAccount_Pin, api.PostAccount_Pin, scope.PostAccount_Pin)
	register(api, e.POST, route.PostAccount_Unpin, api.PostAccount_Unpin, scope.PostAccount_Unpin)
	register(api, e.POST, route.PostAccount_Note, api.PostAccount_Note, scope.PostAccount_Note)
	register(api, e.GET, route.PostAccount_Relationships, api.GetAccount_Relationships, scope.PostAccount_Relationships)
	register(api, e.GET, route.GetAccount_FamiliarFollowers, api.GetAccount_FamiliarFollowers, scope.GetAccount_FamiliarFollowers)
	register(api, e.GET, route.GetAccount_Search, api.GetAccount_Search, scope.GetAccount_Search)
	register(api, e.GET, route.GetAccount_Lookup, api.GetAccount_Lookup, scope.GetAccount_Lookup)

	// https://docs.joinmastodon.org/methods/apps/
	register(api, e.POST, route.PostApplication, api.PostApplication, scope.PostApplication)
	register(api, e.GET, route.GetApplication_VerifyCredentials, api.GetApplication_VerifyCredentials, scope.GetApplication_VerifyCredentials)
}

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
