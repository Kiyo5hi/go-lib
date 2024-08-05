package envs

type Enviroment struct {
	Runtime *Runtime
}

func NewEnviroment() *Enviroment {
	return &Enviroment{
		Runtime: NewRuntime(),
	}
}
