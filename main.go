//go:generate goagen bootstrap -d github.com/made2591/go-blockchain-go/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/made2591/go-blockchain-go/app"
)

func main() {
	// Create service
	service := goa.New("go-blockchain-go")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "block" controller
	c := NewBlockController(service)
	app.MountBlockController(service, c)
	// Mount "health" controller
	c2 := NewHealthController(service)
	app.MountHealthController(service, c2)
	// Mount "swagger" controller
	c3 := NewSwaggerController(service)
	app.MountSwaggerController(service, c3)

	// Start service
	if err := service.ListenAndServe(":8081"); err != nil {
		service.LogError("startup", "err", err)
	}

}
