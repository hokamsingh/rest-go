package upload

import (
	"log"

	LessGo "github.com/hokamsingh/lessgo/pkg/lessgo"
)

type UploadController struct {
	LessGo.BaseController
	Path    string
	Service UploadService
}

func NewUploadController(service *UploadService, path string) *UploadController {
	// if !LessGo.ValidatePath(path) {
	// 	log.Fatalf("Invalid path provided: %s", path)
	// }
	return &UploadController{Path: path, Service: *service}
}

func (uc *UploadController) RegisterRoutes(r *LessGo.Router) {
	ur := r.SubRouter("/upload")
	log.Print(uc.Path)
	ur.Get("/fs", func(ctx *LessGo.Context) {
		ctx.Send("I am fs")
	})

	ur.Post("/files", func(ctx *LessGo.Context) {
		ctx.Send("file saved")
	})
}
