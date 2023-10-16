package main

import (
	"crud_dynamo/config"
	"crud_dynamo/db"
	"crud_dynamo/handlers"
	"crud_dynamo/middlewares"
	"log"
	"os"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func main() {
	cfg := config.LoadConfigFromEnv()
	db.InitDB(cfg)
	handlers.InitHandlers(cfg)

	// Check for secrets
	expectedToken := os.Getenv("X_API_Secret")
	if expectedToken == "" {
		panic("AUTH_TOKEN is not set in the environment")
	}

	// Define the router
	r := router.New()

	// Define the routes
	r.GET("/get_state", middlewares.LoggingMiddleware(middlewares.AuthMiddleware(handlers.GetStateHandler)))
	r.POST("/set_state", middlewares.LoggingMiddleware(middlewares.AuthMiddleware(handlers.PostStateHandler)))
	r.PUT("/update_state", middlewares.LoggingMiddleware(middlewares.AuthMiddleware(handlers.UpdateStateHandler)))
	r.DELETE("/delete_state", middlewares.LoggingMiddleware(middlewares.AuthMiddleware(handlers.DeleteStateHandler)))

	// Start server
	log.Println("Server started on :8080")
	if err := fasthttp.ListenAndServe(":8080", r.Handler); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}
