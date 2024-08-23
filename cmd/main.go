package main

import (
	"log"
	"time"

	"github.com/hokamsingh/lessgo/app/src"
	"github.com/hokamsingh/lessgo/app/src/upload"
	user "github.com/hokamsingh/lessgo/app/src/user"
	LessGo "github.com/hokamsingh/lessgo/pkg/lessgo"
	// Ensure you're using the correct DI package
)

func main() {
	// Load Configuration
	cfg := LessGo.LoadConfig()
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
	folderPath, err := LessGo.GetFolderPath("uploads")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	App.ServeStatic("/static/", folderPath)

	// Dependency Injection Container
	container := LessGo.NewContainer()

	// Register Services
	if err := container.Register(user.NewUserService); err != nil {
		log.Fatalf("Error registering UserService: %v", err)
	}
	if err := container.Register(upload.NewUploadService); err != nil {
		log.Fatalf("Error registering UploadService: %v", err)
	}
	// Register Modules (using module.IModule interface)
	if err := container.Register(user.NewUserModule); err != nil {
		log.Fatalf("Error registering UserModule: %v", err)
	}
	if err := container.Register(upload.NewUploadModule); err != nil {
		log.Fatalf("Error registering UploadModule: %v", err)
	}

	// _, err = LessGo.DiscoverModules()
	// if err != nil {
	// 	log.Fatalf("Error discovering modules: %v", err)
	// }

	// Root Module
	rootModule := src.NewRootModule(App)
	LessGo.RegisterModules(App, []LessGo.IModule{rootModule})

	// Example Route
	App.Get("/ping", func(ctx *LessGo.Context) {
		ctx.Send("pong")
	})

	// Start the server
	serverPort := cfg.Get("SERVER_PORT", "8080")
	env := cfg.Get("ENV", "development")
	log.Printf("Starting server on port %s in %s mode", serverPort, env)
	if err := App.Listen(":" + serverPort); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
