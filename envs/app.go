package envs

import "os"

type AppName string

func NewAppName() AppName {
	raw, ok := os.LookupEnv("APP_NAME")
	if !ok {
		return AppName(os.Args[0])
	}
	return AppName(raw)
}
