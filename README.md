# toot-echo

[![GoDoc](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://pkg.go.dev/github.com/benpate/toot-echo)
[![Version](https://img.shields.io/github/v/release/benpate/toot-echo?include_prereleases&style=flat-square&color=brightgreen)](https://github.com/benpate/toot-echo/releases)
[![Build Status](https://img.shields.io/github/actions/workflow/status/benpate/toot-echo/go.yml?style=flat-square)](https://github.com/benpate/toot-echo/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/benpate/toot-echo?style=flat-square)](https://goreportcard.com/report/github.com/benpate/toot-echo)
[![Codecov](https://img.shields.io/codecov/c/github/benpate/toot-echo.svg?style=flat-square)](https://codecov.io/gh/benpate/toot-echo)


## Echo adapter for the toot Mastodon API

Toot-echo connects the [toot](http://github.com/benpate/toot-echo) library to the [echo](https://echo.labstack.com/) router.  This makes it easy for you to Mastodon-ize your Go/echo application.

```go
api.PostStatus = func(authorization model.Authorization, values txn.PostStatus) (object.Status, error) {
	// Do the thing that:
	// 1) creates a new `status`
	// 2) writes it to the database
	// 3) returns the value to the caller
}

// Once your handlers are defined, connect Toot to your router like this:
e := echo.New()
tootecho.Register(e, api)
```

## Routers

Currently, toot-echo is the only router adapter I have created for Toot.  It should, however, be a simple project to write adapters for other routers.  If you'd like to use Toot with a different router, please get in touch and let's work together.

## Project Status (Alpha)

This project is still in early testing, as I continue filling out the Mastodon support in [Emissary](https://emissary.dev).  Interfaces may change in the future if I run into trouble making things work with this reference implementation.

However, at this point, feedback from other developers will be immensely helpful in making Tooth a useful tool for the Go community.  Please feel free to open issues, or submit pull requests.
