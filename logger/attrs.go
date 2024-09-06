package logger

import (
	"encoding/json"
	"log/slog"
)

var (
	String = slog.String
	Any    = slog.Any
)

func JsonKv[T any](key string, value T) slog.Attr {
	b, _ := json.Marshal(value)
	return String(key, string(b))
}

func ErrorKv(err error) slog.Attr {
	return slog.String("error", err.Error())
}
