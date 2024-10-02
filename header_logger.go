package headerlogger

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

// HeaderLogger is a Caddy module that logs request and response headers.
type HeaderLogger struct{}

// CaddyModule returns the Caddy module information.
func (HeaderLogger) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.header_logger",
		New: func() caddy.Module { return new(HeaderLogger) },
	}
}

// ServeHTTP implements the caddyhttp.MiddlewareHandler interface.
func (h HeaderLogger) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	// Log incoming request headers
	fmt.Printf("Incoming Request Headers:\n")
	for name, values := range r.Header {
		for _, value := range values {
			fmt.Printf("%s: %s\n", name, value)
		}
	}
    
	// Capture response headers and body
	rec := caddyhttp.NewResponseRecorder(w, nil, func(status int, header http.Header) bool {
		return true
	})

	err := next.ServeHTTP(rec, r)
	
	// Log response headers
	fmt.Printf("Outgoing Response Headers:\n")
	for name, values := range rec.Header() {
		for _, value := range values {
			fmt.Printf("%s: %s\n", name, value)
		}
	}

	rec.WriteResponse()
	return err
}

// Interface guard
var _ caddyhttp.MiddlewareHandler = (*HeaderLogger)(nil)

// init registers the module.
func init() {
	caddy.RegisterModule(HeaderLogger{})
}

