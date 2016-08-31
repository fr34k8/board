package context

import (
	"net/http"
	"github.com/op/go-logging"
	"github.com/julienschmidt/httprouter"
	"github.com/mtti/board/server/repositories"
)

// Global context
type Global struct {
	Repository repositories.Repository
	Log *logging.Logger
}

// Wraps any number of middleware into a httprouter handler
func (ctx *Global) WrapRouteHandlers(handlers ...Middleware) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Build request context object
		req := &Request{
			Global: ctx,
			Log: ctx.Log,
			Writer: w,
			Request: r,
			Params: ps,
		}
		// Execute each middleware until one returns true to stop handling
		for _, handler := range handlers {
			if handler(req) {
				break
			}
		}
	}	
}

// Request context
type Request struct {
	Global *Global
	Log *logging.Logger
	Writer http.ResponseWriter
	Request *http.Request
	Params httprouter.Params
}

type Middleware func(*Request) bool
