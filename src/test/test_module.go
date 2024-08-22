package test

import (
	LessGo "github.com/hokamsingh/lessgo/pkg/lessgo"
)

func NewTestModule() *LessGo.Module {
	testService := NewTestService()
	testController := NewTestController(testService, "/test")
	return LessGo.NewModule("ExampleModule",
		[]interface{}{testController}, // Controllers
		[]interface{}{testService},    // Services
	)
}
