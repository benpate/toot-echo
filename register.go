package tootecho

import (
	"net/http"

	"github.com/benpate/derp"
	"github.com/benpate/toot"
	"github.com/labstack/echo/v4"
)

func Register(e *echo.Echo, api toot.API) {

	// https://docs.joinmastodon.org/methods/accounts/
	e.POST("/api/v1/accounts", wrap(api.PostAccounts))
	e.GET("/api/v1/accounts/verify_credentials", wrap(api.GetAccounts_VerifyCredentials))
	e.PATCH("/api/v1/accounts/update_credentials", wrap(api.PatchAccounts_UpdateCredentials))
	e.GET("/api/v1/accounts/:id", wrap(api.GetAccount))
	e.GET("/api/v1/accounts/:id/statuses", wrap(api.GetAccount_Statuses))
	e.GET("/api/v1/accounts/:id/followers", wrap(api.GetAccount_Followers))
	e.GET("/api/v1/accounts/:id/following", wrap(api.GetAccount_Following))
	e.GET("/api/v1/accounts/:id/featured_tags", wrap(api.GetAccount_FeaturedTags))
	e.POST("/api/v1/accounts/:id/follow", wrap(api.PostAccount_Follow))
	e.POST("/api/v1/accounts/:id/unfollow", wrap(api.PostAccount_Unfollow))
	e.POST("/api/v1/accounts/:id/block", wrap(api.PostAccount_Block))
	e.POST("/api/v1/accounts/:id/unblock", wrap(api.PostAccount_Unblock))
	e.POST("/api/v1/accounts/:id/mute", wrap(api.PostAccount_Mute))
	e.POST("/api/v1/accounts/:id/unmute", wrap(api.PostAccount_Unmute))
	e.POST("/api/v1/accounts/:id/pin", wrap(api.PostAccount_Pin))
	e.POST("/api/v1/accounts/:id/unpin", wrap(api.PostAccount_Unpin))
	e.POST("/api/v1/accounts/:id/note", wrap(api.PostAccount_Note))
	e.GET("/api/v1/accounts/relationships", wrap(api.GetAccount_Relationships))
	e.GET("/api/v1/accounts/:id/familiar_followers", wrap(api.GetAccount_FamiliarFollowers))
	e.GET("/api/v1/accounts/search", wrap(api.GetAccount_Search))
	e.GET("/api/v1/accounts/lookup", wrap(api.GetAccount_Lookup))

	// https://docs.joinmastodon.org/methods/apps/
	e.POST("/api/v1/apps", wrap(api.PostApplication))
	e.GET("/api/v1/apps/verify_credentials", wrap(api.GetApplication_VerifyCredentials), authorize(api))

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
