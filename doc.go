// Package tootecho serves a Mastodon API on an Echo web server.
//
// Fill in the handlers of a toot.API and hand it to Register, which walks every
// endpoint the toot library defines and adds a route for each handler that is
// present:
//
//	api := toot.New(myAuthorizer)
//	api.GetAccount = myGetAccountHandler
//	tootecho.Register(e, api)
//
// A nil handler is skipped rather than registered, so an API that implements ten
// endpoints serves ten routes and 404s the rest. Middleware passed to Register
// runs on every route it adds.
//
// Each route wraps its handler in the same three steps: authorize the request and
// check its required OAuth scope, bind the request into the handler's input struct
// from the txn package, and marshal whatever the handler returns as JSON.
//
// # What the host application still owns
//
//   - Error rendering. Handlers report failure by returning a derp error, and this
//     package passes it up to Echo unchanged. Echo's default handler renders any
//     error it does not recognize as a 500, so install a derp-aware
//     HTTPErrorHandler to get real status codes -- 401 for a rejected token, 404
//     for a missing record, and so on.
//   - Trusting the proxy. The WithHost middleware prefers the X-Forwarded-Host
//     header over the request's own Host, because a Mastodon API has to report the
//     public hostname of the instance. That header is caller-supplied: run this
//     behind a proxy that overwrites it, or a client can choose the hostname that
//     appears in the responses it receives.
//
// Every response carries "Access-Control-Allow-Origin: *", matching Mastodon, and
// safe here because authorization rides on a Bearer token rather than on a cookie.
package tootecho
