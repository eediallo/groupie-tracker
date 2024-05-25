package logstack

import (
	"log/slog"
	"os"
)

// ----------------------------------------
func loadLoggerToCli() error {
	funcName := "loadLoggerToCli"

	//----------------------
	logHandlerOpts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
		//AddSource: true,
	}
	logHandler := slog.NewTextHandler(os.Stderr, logHandlerOpts)
	loggerToCli = slog.New(logHandler)
	slog.Info(LogMsg(pkgName, funcName, "loadLoggerToCli", "loading logger to Cli stack", nil))
	return nil
}

// ----------------------------------------
