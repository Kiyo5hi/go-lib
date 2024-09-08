package logger

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	slogmulti "github.com/samber/slog-multi"
)

func Logger(opts ...option) *slog.Logger {
	cfg := Options(opts...)
	var handler slog.Handler
	if cfg.debug {
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelDebug,
		})
	} else {
		reader, _ := os.Open(fmt.Sprintf("%s.log", string(cfg.appName)))
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
