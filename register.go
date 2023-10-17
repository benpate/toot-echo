package tootecho

import (
	"fmt"
	"net/http"

	"github.com/benpate/derp"
	"github.com/benpate/toot"
	"github.com/benpate/toot/route"
	"github.com/labstack/echo/v4"
)

type echoMethod func(string, echo.HandlerFunc, ...echo.MiddlewareFunc) *echo.Route

func Register(e *echo.Echo, api toot.API) {

	fmt.Println("REGISTER: BEGIN")

	// https://docs.joinmastodon.org/methods/accounts/
	register(e.POST, route.PostAccount, api.PostAccount)
	register(e.GET, route.GetAccount_VerifyCredentials, api.GetAccount_VerifyCredentials)
	register(e.PATCH, route.PatchAccount_UpdateCredentials, api.PatchAccount_UpdateCredentials)
	register(e.GET, route.GetAccount, api.GetAccount)
	register(e.GET, route.GetAccount_Statuses, api.GetAccount_Statuses)
	register(e.GET, route.GetAccount_Followers, api.GetAccount_Followers)
	register(e.GET, route.GetAccount_Following, api.GetAccount_Following)
	register(e.GET, route.GetAccount_FeaturedTags, api.GetAccount_FeaturedTags)
	register(e.POST, route.PostAccount, api.PostAccount_Follow)
	register(e.POST, route.PostAccount_Unfollow, api.PostAccount_Unfollow)
	register(e.POST, route.PostAccount_Block, api.PostAccount_Block)
	register(e.POST, route.PostAccount_Unblock, api.PostAccount_Unblock)
	register(e.POST, route.PostAccount_Mute, api.PostAccount_Mute)
	register(e.POST, route.PostAccount_Unmute, api.PostAccount_Unmute)
	register(e.POST, route.PostAccount_Pin, api.PostAccount_Pin)
	register(e.POST, route.PostAccount_Unpin, api.PostAccount_Unpin)
	register(e.POST, route.PostAccount_Note, api.PostAccount_Note)
	register(e.GET, route.PostAccount_Relationships, api.GetAccount_Relationships)
	register(e.GET, route.GetAccount_FamiliarFollowers, api.GetAccount_FamiliarFollowers)
	register(e.GET, route.GetAccount_Search, api.GetAccount_Search)
	register(e.GET, route.GetAccount_Lookup, api.GetAccount_Lookup)

	// https://docs.joinmastodon.org/methods/apps/
	register(e.POST, route.PostApplication, api.PostApplication)
	register(e.GET, route.GetApplication_VerifyCredentials, api.GetApplication_VerifyCredentials, authorize(api))
}

func register[Input any, Output any](method echoMethod, path string, handler toot.APIFunc[Input, Output], middleware ...echo.MiddlewareFunc) {
	if handler != nil {
		fmt.Println("REGISTER-YUP: " + path)
		method(path, wrap(handler), middleware...)
		return
	}

	fmt.Println("REGISTER-NAH: " + path)
}

func authorize(api toot.API, scopes ...string) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {

		return func(ctx echo.Context) error {

			if api.Authorize != nil {
				if !api.Authorize(ctx.Request(), scopes...) {
					return derp.NewForbiddenError("toot-echo.verifyScopes", "You do not have permission to access this resource", scopes)
				}
			}

			return next(ctx)
		}
	}
}

func wrap[In any, Out any](handler toot.APIFunc[In, Out]) echo.HandlerFunc {

	return func(ctx echo.Context) error {

		if handler == nil {
			return derp.NewInternalError("toot-echo.wrap", "This API is not implemented")
		}

		var args In

		// Collect input arguments from the Request
		// TODO: Replace this bind with custom binder:
		// https://github.com/go-playground/form
		// https://echo.labstack.com/docs/binding#custom-binding
		if err := ctx.Bind(&args); err != nil {
			return derp.Wrap(err, "toot-echo.wrap", "Error parsing request body")
		}

		// Execute the API handler
		result, err := handler(ctx.Request(), args)

		if err != nil {
			return derp.Wrap(err, "toot-echo.wrap", "Error executing API call")
		}

		// Return the API result to the caller as JSON
		if err := ctx.JSON(http.StatusOK, result); err != nil {
			return derp.Wrap(err, "toot-echo.wrap", "Error writing response body")
		}

		// Woot.
		return nil
	}
}
