// main.go

// @title My API
// @version 1.0
// @description This is a sample server.
// @host localhost:8080
// @BasePath /

package main

import (
	"log"
	"time"

	"github.com/hokamsingh/lessgo/app/src"
	LessGo "github.com/hokamsingh/lessgo/pkg/lessgo"
)

func main() {
	// Load Configuration
	cfg := LessGo.LoadConfig()
	serverPort := cfg.Get("SERVER_PORT", "8080")
	env := cfg.Get("ENV", "development")
	addr := ":" + serverPort

	// CORS Options
	corsOptions := LessGo.NewCorsOptions(
		[]string{"*"}, // Allow all origins
		[]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Allowed methods
		[]string{"Content-Type", "Authorization"},           // Allowed headers
	)

	// Initialize App
	App := LessGo.App(
		LessGo.WithCORS(*corsOptions),
		LessGo.WithRateLimiter(100, 1*time.Minute),
		LessGo.WithJSONParser(),
		LessGo.WithCookieParser(),
		// LessGo.WithFileUpload("uploads"),
	)

	// Serve Static Files
	folderPath, _ := LessGo.GetFolderPath("uploads")
	App.ServeStatic("/static/", folderPath)

	// Register dependencies
	dependencies := []interface{}{src.NewRootService, src.NewRootModule}
	LessGo.RegisterDependencies(dependencies)

	// Root Module
	rootModule := src.NewRootModule(App)
	LessGo.RegisterModules(App, []LessGo.IModule{rootModule})

	// Example Route
	App.Get("/ping", func(ctx *LessGo.Context) {
		ctx.Send("pong")
	})

	// Start the server
	log.Printf("Starting server on port %s in %s mode", serverPort, env)
	if err := App.Listen(addr); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
