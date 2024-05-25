package logstack

import (
	//"log"
	"log/slog"
	"os"
	"time"
)

// -------------------------------------------
func loadLoggerToFile() error {
	//-------------------------------
	funcName := "loadLoggerToFile"
	var opName, opDes string
	var err error
	//--------------------------------
	todayDate := time.Now().Format("2006-01-02")
	logFileAddress := LogFilesDirectory + "log-" + todayDate + ".json"
	//-------------------------------
	var logFile *os.File
	logFile, err = os.OpenFile(logFileAddress, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	opName, opDes = "OpenFile", "logger file I/O"
	if err != nil {
		return RSlogErr(pkgName, funcName, opName, opDes, err)
	} else {
		slog.Info(LogMsg(pkgName, funcName, opName, opDes, nil))
	}
	//-------------------------------
	logHandlerOptsFile := &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}
	loggerToFile = slog.New(slog.NewJSONHandler(logFile, logHandlerOptsFile))
	opName, opDes = "slog.New", "slog logger To File"
	slog.Info(LogMsg(pkgName, funcName, opName, opDes, nil))
	//-------------------------------
	return nil
}

//-------------------------
