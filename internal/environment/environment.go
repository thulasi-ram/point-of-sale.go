package environment

import "os"

type Environment string

const (
	DevEnv        Environment = "development"
	TestingEnv    Environment = "test"
	StagingEnv    Environment = "staging"
	ProductionEnv Environment = "production"
)

type EnvProvider interface {
	CurrEnv() Environment
}

type osEnv struct {
	key string
}

func NewOsEnv() *osEnv {
	return &osEnv{key: "ENV"}
}

func (e *osEnv) CurrEnv() Environment {
	env := Environment(os.Getenv(e.key))
	switch env {
	case DevEnv, TestingEnv, StagingEnv, ProductionEnv:
		return env
	default:
		panic("current os env is not a valid value")
	}
}

func (e Environment) IsStructuredLogging() bool {
	switch e {
	case ProductionEnv, StagingEnv:
		return true
	default:
		return false
	}
}
