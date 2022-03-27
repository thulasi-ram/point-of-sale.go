package web

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-logr/logr"
	"github.com/jackc/pgx/v4/pgxpool"
	"net/http"
	"os"
	"point-of-sale.go/v1/internal/config"
	"point-of-sale.go/v1/internal/environment"
	"point-of-sale.go/v1/internal/logger"
)

type App struct {
	env    environment.Environment
	log    logr.Logger
	config config.Config
	db     *pgxpool.Pool

	router chi.Router
}

func NewApp() *App {
	env := environment.NewOsEnv().CurrEnv()
	log := logger.NewLogr(env)
	cfg := config.NewConfig(env)

	db, err := pgxpool.Connect(context.Background(), cfg.DATABASE_URL)
	if err != nil {
		fmt.Println("Unable to connect to db")
		os.Exit(1)
	}

	router := chi.NewRouter()

	return &App{env: env, log: log, config: cfg, router: router, db: db}
}

func (a *App) CurrEnv() environment.Environment {
	return a.env
}

func (a *App) Logger() logr.Logger {
	return a.log
}

func (a *App) LoggerFromContext(ctx context.Context) logr.Logger {
	l := logger.WithRequestID(ctx, a.log)
	return l
}

func (a *App) Config() config.Config {
	return a.config
}

func (a *App) Router() chi.Router {
	return a.router
}

func (a *App) DB() *pgxpool.Pool {
	return a.db
}

func (a *App) WithMiddlewares(middlewares ...func(handler http.Handler) http.Handler) *App {
	for _, mdw := range middlewares {
		a.router.Use(mdw)
	}
	return a
}

func (a *App) Close() {
	a.db.Close()
}
