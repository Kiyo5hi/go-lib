package logger

import (
	"context"
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

func Debug(ctx context.Context, msg string, attrs ...slog.Attr) {
	slog.LogAttrs(ctx, slog.LevelDebug, msg, attrs...)
}

func Info(ctx context.Context, msg string, attrs ...slog.Attr) {
	slog.LogAttrs(ctx, slog.LevelInfo, msg, attrs...)
}

func Warn(ctx context.Context, msg string, attrs ...slog.Attr) {
	slog.LogAttrs(ctx, slog.LevelWarn, msg, attrs...)
}

func Error(ctx context.Context, msg string, attrs ...slog.Attr) {
	slog.LogAttrs(ctx, slog.LevelError, msg, attrs...)
}

func Fatal(ctx context.Context, msg string, attrs ...slog.Attr) {
	slog.LogAttrs(ctx, 12, msg, attrs...)
	os.Exit(1)
}
