package logger

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/kiyo5hi/go-lib/envs"
	slogmulti "github.com/samber/slog-multi"
)

func Logger(opts ...option) *slog.Logger {
	cfg := Options(opts...)
	var handler slog.Handler
	environ := envs.NewEnviroment()
	if cfg.debug {
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
