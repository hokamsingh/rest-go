package main

import (
	"log"
	"time"

	"github.com/hokamsingh/lessgo/app/src"
	user "github.com/hokamsingh/lessgo/app/src/users"
	LessGo "github.com/hokamsingh/lessgo/pkg/lessgo"
)

func main() {
	// Load Configuration
	cfg := LessGo.LoadConfig()

	// Cors
	corsOptions := LessGo.NewCorsOptions(
		[]string{"*"}, // Allow all origins
		[]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Allowed methods
		[]string{"Content-Type", "Authorization"},           // Allowed headers
	)

	app := LessGo.App(
		LessGo.WithCORS(*corsOptions),
		LessGo.WithRateLimiter(100, 1*time.Minute),
		LessGo.WithJSONParser(),
		LessGo.WithCookieParser(),
	)

	// SERVE STATIC
	folderPath, err := LessGo.GetFolderPath("uploads")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	app.ServeStatic("/static/", folderPath)

	container := LessGo.NewContainer()
	if err := container.Register(user.NewUserService); err != nil {
		log.Fatalf("Error registering UserService: %v", err)
	}
	// if err := container.Register(upload.NewUploadService); err != nil {
	// 	log.Fatalf("Error registering UploadService: %v", err)
	// }
	// if err := container.Register(src.NewRootService); err != nil {
	// 	log.Fatalf("Error registering RootService: %v", err)
	// }

	// Optionally register modules if they are separate from services
	if err := container.Register(user.NewUserModule); err != nil {
		log.Fatalf("Error registering UserModule: %v", err)
	}
	// if err := container.Register(upload.NewUploadModule); err != nil {
	// 	log.Fatalf("Error registering UploadModule: %v", err)
	// }
	// if err := container.Register(src.NewRootModule); err != nil {
	// 	log.Fatalf("Error registering RootModule: %v", err)
	// }

	// // Register services and modules in the container
	// // container.Register(test.NewTestService)
	// container.Register(user.NewUserService)
	// container.Register(upload.NewUploadService)
	// container.Register(src.NewRootService)

	// // container.Register(test.NewTestModule)
	// container.Register(user.NewUserModule)
	// container.Register(upload.NewUploadModule)
	// container.Register(src.NewRootModule)

	// Invoke and register Test Module routes
	// Register routes for TestModule
	rootModule := src.NewRootModule(app, container)
	LessGo.RegisterModules(app, container, []LessGo.Module{*rootModule})
	// testModule := test.NewTestModule()
	// uploadModule := upload.NewUploadModule()
	// modules := []LessGo.Module{*testModule, *uploadModule}
	// LessGo.RegisterModules(r, container, modules)

	// apiRouter := r.SubRouter("/api")

	// apiRouter.Get("/test", func(ctx *LessGo.Context) {
	// 	ctx.Send("success")
	// })

	// Define a custom route with new context
	app.Get("/ping", func(ctx *LessGo.Context) {
		// ctx.JSON(200, map[string]string{"message": "pong"})
		ctx.Send("pong")
	})

	// r.Get("/user/{id}", func(ctx *LessGo.Context) {
	// 	// Get all URL params
	// 	params, ok := ctx.GetAllParams()
	// 	id := params["id"]
	// 	if !ok {
	// 		ctx.Error(400, "no params found")
	// 		return
	// 	}
	// 	// Get all query params as JSON
	// 	queryParams, _ := ctx.GetAllQuery()
	// 	// Set a custom header
	// 	ctx.SetHeader("X-Custom-Header", "MyValue")
	// 	cookie, ok := ctx.GetCookie("auth_token")
	// 	if !ok {
	// 		// ctx.Error(400, "Bad Request")
	// 		ctx.SetCookie("auth_token", "0xc000013a", 60, "")
	// 	}
	// 	ctx.JSON(200, map[string]interface{}{
	// 		"params":      params,
	// 		"queryParams": queryParams,
	// 		"id":          id,
	// 		"cookie":      cookie,
	// 	})
	// })

	// r.Post("/submit", func(ctx *LessGo.Context) {
	// 	var body test.User
	// 	ctx.Body(&body)
	// 	ctx.JSON(200, body)
	// })

	// r.Delete("/{id}", func(ctx *LessGo.Context) {
	// 	var id string
	// 	id, _ = ctx.GetParam("id")
	// 	ctx.Error(400, id)
	// })

	// Start the server
	serverPort := cfg.Get("SERVER_PORT", "8080")
	env := cfg.Get("ENV", "development")
	log.Printf("Starting server on port %s in %s mode", serverPort, env)
	if err := app.Listen(":" + serverPort); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
