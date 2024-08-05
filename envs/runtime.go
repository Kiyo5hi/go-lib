package envs

import (
	"os"

	"github.com/samber/lo"
)

type Runtime string

const (
	RuntimeDebug      Runtime = "debug"
	RuntimeProduction Runtime = "prod"
)

func NewRuntime() *Runtime {
	raw := os.Getenv("RUNTIME")
	rt := Runtime(raw)
	if rt == RuntimeProduction {
		return lo.ToPtr(RuntimeProduction)
	}
	return lo.ToPtr(RuntimeDebug)
}
