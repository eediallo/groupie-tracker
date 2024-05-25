package logstack

import (
	"errors"
	"log/slog"
)

// -------------------------------
func LoadLoggerStack() error {
	//---------------------------------
	funcName := "LoadLoggerStack"
	var opName, opDes string
	//----------------loadLoggerToFile-----------------

	opName, opDes = "loadLoggerToFile", "loading logger to File stack"
	if err := loadLoggerToFile(); err != nil {
		return RSlogErr(pkgName, funcName, opName, opDes, err)
	} else {
		slog.Info(LogMsg(pkgName, funcName, opName, opDes, err))
	}
	//----------------loadLoggerToFile-----------------
	opName, opDes = "loadLoggerToCli", "loading logger to Cli stack"
	if err := loadLoggerToCli(); err != nil {
		return RSlogErr(pkgName, funcName, opName, opDes, err)
	} else {
		slog.Info(LogMsg(pkgName, funcName, opName, opDes, err))
	}
	//---------------------------------
	slog.Info(LogMsg(pkgName, funcName, "LoadLoggerStack", "loading whole logger stack", nil))
	//---------------------------------
	logger.Info("logStack", "Info", "test Info description")
	logger.Error("logStack", "Error", errors.New("Error"), "test Error description")
	logger.RErr("logStack", "RErr", errors.New("RErr"), "test RErr description")
	logger.Warn("logStack", "Warn", errors.New("Warn"), "test Warn description")
	logger.RWarn("logStack", "RWarn", errors.New("RErr"), "test RWarn description")
	return nil
}
