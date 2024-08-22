package src

import (
	"github.com/hokamsingh/lessgo/app/src/test"
	"github.com/hokamsingh/lessgo/app/src/upload"
	user "github.com/hokamsingh/lessgo/app/src/users"
	LessGo "github.com/hokamsingh/lessgo/pkg/lessgo"
)

func NewRootModule(r *LessGo.Router, c *LessGo.Container) *LessGo.Module {
	// Initialize and collect all modules
	modules := []LessGo.Module{
		*test.NewTestModule(),
		*upload.NewUploadModule(),
		*user.NewUserModule(),
	}

	// Register all modules
	LessGo.RegisterModules(r, c, modules)
	// for _, mod := range modules {
	// 	err := LessGo.RegisterModuleRoutes(r, c, mod)
	// 	if err != nil {
	// 		log.Fatalf("Error registering Module routes: %v", err)
	// 	}
	// }

	service := NewRootService()
	controller := NewRootController(service, "/")
	return LessGo.NewModule("root", []interface{}{controller}, []interface{}{service})
}
