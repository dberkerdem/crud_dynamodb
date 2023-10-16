package middlewares

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"time"
)

func LoggingMiddleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		start := time.Now()

		defer func() {
			fmt.Printf("[%s] %s %s (took %s)\n", ctx.Response.Header.Peek("X-Request-ID"), ctx.Method(), ctx.RequestURI(), time.Since(start))
		}()

		next(ctx)
	}
}
