package web

import (
	"net/http"
	"reflect"
	"runtime"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

// NewRouter creates a new HTTP router and registers the routes.
func NewRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.Handle("/", traceMiddleware(http.HandlerFunc(HelloWorld)))
	return router
}

// traceMiddleware はOpenTelemetryのトレースを行うためのミドルウェアです
func traceMiddleware(handler http.Handler) http.Handler {
	handlerName := runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name()
	return otelhttp.NewHandler(handler, handlerName)
}
