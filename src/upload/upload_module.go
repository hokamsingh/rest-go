package upload

import (
	LessGo "github.com/hokamsingh/lessgo/pkg/lessgo"
)

func NewUploadModule() *LessGo.Module {
	service := NewUploadService("uploads")
	controller := NewUploadController(service, "/upload")
	return LessGo.NewModule("UploadModule",
		[]interface{}{controller},
		[]interface{}{service},
	)
}
