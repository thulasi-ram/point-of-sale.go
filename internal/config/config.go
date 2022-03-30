package config

import (
	"fmt"
	"os"
	"point-of-sale.go/v1/internal/environment"
)

type Config struct {
	DATABASE_URL string
}

func newDevConfig() Config {
	return Config{
		DATABASE_URL: os.Getenv("DATABASE_URL"),
	}
}

func newTestConfig() Config {
	return Config{
		DATABASE_URL: os.Getenv("DATABASE_URL"),
	}
}

func newProdConfig() Config {
	return Config{
		DATABASE_URL: os.Getenv("DATABASE_URL"),
	}
}

func NewConfig(env environment.Environment) Config {
	switch env {
	case environment.DevEnv:
		return newDevConfig()
	case environment.TestingEnv:
		return newTestConfig()
	case environment.ProductionEnv:
		return newProdConfig()
	default:
		panic(fmt.Sprintf("unrecognized environment for Config: %s", env))
	}
}
