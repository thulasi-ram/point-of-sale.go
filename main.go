package main

import (
	"context"
	"encoding/json"
	"fmt"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-logr/logr"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"os/signal"
	"point-of-sale.go/v1/internal/middleware"
	"point-of-sale.go/v1/internal/web"
	"point-of-sale.go/v1/products"
	"point-of-sale.go/v1/purchases"
	"time"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Unable to load .env file")
		os.Exit(1)
	}

	app := web.NewApp()
	app = app.WithMiddlewares(
		middleware.RequestID,
		chimiddleware.RealIP,
		chimiddleware.RequestLogger(middleware.NewRequestLogFormatter(app.Logger())),
		chimiddleware.Recoverer,
	)

	app.Router().Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		err := json.NewEncoder(w).Encode(map[string]string{"message": "pong"})
		if err != nil {
			panic(err)
		}
	})

	products.InitRoutes(app)
	purchases.InitRoutes(app)

	s := &http.Server{
		Addr: "0.0.0.0:8080",
		// set timeouts to avoid slowloris attacks
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      app.Router(),
	}

	gracefulShutDown(app.Logger(), s)

	os.Exit(0)
}

func gracefulShutDown(log logr.Logger, s *http.Server) {
	go func() {
		// start server in go routine and then block on interrupt signal
		if err := s.ListenAndServe(); err != nil {
			log.Error(err, "unable to start server")
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	err := s.Shutdown(ctx)

	if err != nil {
		log.Error(err, "failed to shutdown gracefully")
		os.Exit(1)
	}

	log.Info("shutting down server")
}
