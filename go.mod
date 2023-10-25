module github.com/benpate/toot-echo

go 1.20

replace github.com/benpate/toot => ../toot

require github.com/labstack/echo/v4 v4.11.1

require (
	github.com/benpate/derp v0.28.1
	github.com/benpate/rosetta v0.18.2
	github.com/benpate/toot v0.0.1
)

require (
	github.com/labstack/gommon v0.4.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	golang.org/x/crypto v0.14.0 // indirect
	golang.org/x/net v0.16.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
)
