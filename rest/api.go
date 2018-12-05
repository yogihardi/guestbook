package rest

import (
	"net"
	"net/http"
	"time"

	"github.com/yogihardi/guestbook/service"

	"github.com/yogihardi/guestbook/rest/app"
	"github.com/yogihardi/guestbook/rest/controller"

	"golang.org/x/net/context"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/logging/log15"
	"github.com/goadesign/goa/middleware"
	"github.com/inconshreveable/log15"
	"github.com/sirupsen/logrus"
	"github.com/tylerb/graceful"
)

// Run ..
func Run(ctx context.Context, listener net.Listener, appService service.Service, log log15.Logger) error {
	var err error

	// Create service
	service := goa.New("article-recommendation-worker")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())
	service.WithLogger(goalog15.New(log))

	// Mount "swagger" controller
	c2 := controller.NewSwaggerController(service)
	app.MountSwaggerController(service, c2)
	// Mount "version" controller
	c3 := controller.NewVersionController(service)
	app.MountVersionController(service, c3)

	c1 := controller.NewGuestbookController(service, appService)
	app.MountGuestbookController(service, c1)

	// Setup graceful server
	server := &graceful.Server{
		NoSignalHandling: true,
		Server: &http.Server{
			Handler: service.Mux,
		},
	}

	c := make(chan error, 1)
	go func() {
		c <- server.Serve(listener)
	}()

	select {
	case <-ctx.Done():
		// stoping all workers
		// appService.StopWorkers()

		server.Stop(time.Duration(3) * time.Second)
		<-server.StopChan()
		// draining the channel
		<-c
	case err := <-c:
		if err != nil {
			logrus.Error(err.Error())
		}
	}

	return err
}
