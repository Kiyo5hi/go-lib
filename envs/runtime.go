package envs

import (
	"os"
)

type Runtime string

const (
	RuntimeDebug      Runtime = "debug"
	RuntimeProduction Runtime = "prod"
)

func NewRuntime() Runtime {
	raw, ok := os.LookupEnv("RUNTIME")
	if !ok {
		return RuntimeDebug
	}
	rt := Runtime(raw)
	if rt == RuntimeProduction {
		return RuntimeProduction
	}
	return RuntimeDebug
}
