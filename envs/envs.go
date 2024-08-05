package envs

type Enviroment struct {
	AppName AppName
	Runtime Runtime
}

func NewEnviroment() *Enviroment {
	return &Enviroment{
		AppName: NewAppName(),
		Runtime: NewRuntime(),
	}
}
