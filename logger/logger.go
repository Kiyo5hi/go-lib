package logger

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/kiyo5hi/go-lib/envs"
	slogmulti "github.com/samber/slog-multi"
)

func Logger() *slog.Logger {
	var handler slog.Handler
	environ := envs.NewEnviroment()
	if environ.Runtime == envs.RuntimeDebug {
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelDebug,
		})
	} else {
		reader, _ := os.Open(fmt.Sprintf("%s.log", string(environ.AppName)))
		handler = slogmulti.Fanout(
			slog.NewJSONHandler(reader, &slog.HandlerOptions{
				AddSource: true,
				Level:     slog.LevelInfo,
			}),
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
				AddSource: true,
				Level:     slog.LevelInfo,
			}),
		)
	}
	return slog.New(handler)
}
