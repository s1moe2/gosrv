package server

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/s1moe2/gosrv/config"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type apiServer struct {
	httpServer *http.Server
}

func newServer(serverConfig config.ServerConfig, router *mux.Router) *apiServer {
	return &apiServer{
		httpServer: &http.Server{
			Addr: serverConfig.Address,
			//ErrorLog:     log.New(logrus.New().Writer(), "", 0),
			Handler:      http.TimeoutHandler(router, 10*time.Second, "request timeout"),
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 15 * time.Second,
			IdleTimeout:  20 * time.Second,
		},
	}
}

func (s *apiServer) start() error {
	//channel to listen for errors coming from the listener.
	serverErrors := make(chan error, 1)

	go func() {
		log.Printf("main : API listening on %s", s.httpServer.Addr)
		serverErrors <- s.httpServer.ListenAndServe()
	}()

	// channel to listen for an interrupt or terminate signal from the OS.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// blocking run and waiting for shutdown.
	select {
	case err := <-serverErrors:
		return fmt.Errorf("error: starting server: %s", err)

	case <-shutdown:
		log.Println("main : Start shutdown")

		// give outstanding requests a deadline for completion.
		const timeout = 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		// asking listener to shutdown
		err := s.httpServer.Shutdown(ctx)
		if err != nil {
			log.Printf("main : Graceful shutdown did not complete in %v : %v", timeout, err)
			err = s.httpServer.Close()
		}

		if err != nil {
			return fmt.Errorf("main : could not stop server gracefully : %v", err)
		}
	}

	return nil
}
