package web

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-logr/logr"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgtype"
	shopspring "github.com/jackc/pgtype/ext/shopspring-numeric"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/vmihailenco/msgpack/v5"
	"net/http"
	"os"
	"point-of-sale.go/v1/internal/config"
	"point-of-sale.go/v1/internal/environment"
	"point-of-sale.go/v1/internal/logger"
	"strings"
)

type App struct {
	env       environment.Environment
	log       logr.Logger
	config    config.Config
	db        *pgxpool.Pool
	router    chi.Router
	validator *validator.Validate
}

func NewApp() *App {
	env := environment.NewOsEnv().CurrEnv()
	log := logger.NewLogr(env)
	cfg := config.NewConfig(env)
	vldtr := validator.New()

	pgCfg, err := pgxpool.ParseConfig(cfg.DATABASE_URL)
	if err != nil {
		fmt.Println("Unable to parse db url")
		os.Exit(1)
	}

	pgCfg.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		conn.ConnInfo().RegisterDataType(pgtype.DataType{
			Value: &shopspring.Numeric{},
			Name:  "numeric",
			OID:   pgtype.NumericOID,
		})
		return nil
	}

	db, err := pgxpool.ConnectConfig(context.Background(), pgCfg)
	if err != nil {
		fmt.Println("Unable to connect to db")
		os.Exit(1)
	}

	router := chi.NewRouter()

	return &App{env: env, log: log, config: cfg, router: router, db: db, validator: vldtr}
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

func (a *App) Validator() *validator.Validate {
	return a.validator
}

func (a *App) WithMiddlewares(middlewares ...func(handler http.Handler) http.Handler) *App {
	for _, mdw := range middlewares {
		a.router.Use(mdw)
	}
	return a
}

func (a *App) Validate(v interface{}, r *http.Request, body, query, path, header bool) error {
	err := a.Bind(v, r, body, query, path, header)
	if err != nil {
		return err
	}

	err = a.Validator().Struct(v)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) Bind(v interface{}, r *http.Request, body, query, path, header bool) error {

	if body {
		d := json.NewDecoder(r.Body)
		d.DisallowUnknownFields() // safety check
		err := d.Decode(v)
		if err != nil {
			return err
		}
	}

	if query {
		var queryMap map[string]interface{}
		for key, val := range r.URL.Query() {
			queryMap[key] = strings.Join(val, "|")
		}
		if len(queryMap) > 0 {
			err := marshUnmarshMsgpack(v, queryMap)
			if err != nil {
				return err
			}
		}
	}

	if path {
		pathMap := make(map[string]interface{})
		rc := chi.RouteContext(r.Context()).URLParams
		for i := 0; i < len(rc.Keys); i++ {
			pathMap[rc.Keys[i]] = rc.Values[i]
		}
		if len(pathMap) > 0 {
			err := marshUnmarshMsgpack(v, pathMap)
			if err != nil {
				return err
			}
		}
	}

	if header {
		var headerMap map[string]interface{}
		for key, val := range r.Header {
			headerMap[key] = strings.Join(val, "|")
		}
		if len(headerMap) > 0 {
			err := marshUnmarshMsgpack(v, headerMap)
			if err != nil {
				return err
			}
		}

	}

	return nil
}

func (a *App) Close() {
	a.db.Close()
}

func marshUnmarsh(v interface{}, data map[string]interface{}) error {
	// for stores other than streams eg: not req.Body
	// marshall and unmarshall is faster and more efficient
	dataStr, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = json.Unmarshal(dataStr, v)
	if err != nil {
		return err
	}
	return nil
}

func marshUnmarshMsgpack(v interface{}, data map[string]interface{}) error {
	b, err := msgpack.Marshal(data)
	if err != nil {
		return err
	}
	err = msgpack.Unmarshal(b, v)
	if err != nil {
		return err
	}
	return nil
}
