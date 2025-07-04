package logging

import (
	"log/slog"
	"os"
)

var (
	Log     *slog.Logger
	level   = new(slog.LevelVar)
	handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level})
)

func init() {
	Log = slog.New(handler)
	level.Set(slog.LevelInfo)
}

func SetLevel(lvl slog.Level) {
	level.Set(lvl)
}
