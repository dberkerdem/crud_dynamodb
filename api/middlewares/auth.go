package middlewares

import (
	"crud_dynamo/utils"
	"log"
	"os"

	"github.com/valyala/fasthttp"
)

func AuthMiddleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		expectedToken := os.Getenv("X_API_Secret")
		authHeader := ctx.Request.Header.Peek("X-API-Secret")

		if authHeader == nil || string(authHeader) != expectedToken {
			// Log the unauthorized access attempt
			log.Printf("Unauthorized access attempt. Received token: %s, Expected token: %s\n", string(authHeader), expectedToken)
			utils.ErrorHandler(ctx, utils.NewHTTPError(fasthttp.StatusUnauthorized, "Unauthorized"))
			return
		}

		next(ctx)
	}
}
