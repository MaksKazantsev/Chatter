package log

import (
	"context"
	"log/slog"
	"os"
)

type Wrapper func(ctx context.Context) context.Context

const LoggerKey = "loggerKey"

type Logger struct {
	*slog.Logger
}

func InitLogger(env string) Logger {
	var log *slog.Logger
	switch env {
	case "local":
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case "prod":
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))

	default:
		panic("unknown env: " + env)
	}
	return Logger{log}
}

func WithLogger(ctx context.Context, log Logger) context.Context {
	return context.WithValue(ctx, LoggerKey, log)

}
func GetLogger(ctx context.Context) Logger {
	return ctx.Value(LoggerKey).(Logger)
}
