package logstack

import (
	"errors"
	"log/slog"
)

// --------------const-----------------------------

const (
	LogFilesDirectory = "logs/"
	pkgName           = "logStack"
)

// -----------var----------------------
var (
	loggerToFile *slog.Logger
	loggerToCli  *slog.Logger

	logger LogCollector

	//funcsAccessList  map[string]map[string]map[string]bool
)

// ===================struct=========================
type LogCollector struct {
	PackageName string
}

// ---------------------------------
func (l LogCollector) ErrMsg(FuncName string, OperationName string, RetunedError error) error {
	return errMsg(FuncName, OperationName, RetunedError)
}

// ---------------------------------
func (l LogCollector) Info(FuncName string, OperationName string, operationDescription string) {

	operationDescription = " success: " + operationDescription

	//loggerToCli.Info(LogMsg(l.PackageName, FuncName, OperationName, operationDescription, nil))
	loggerToFile.Info(LogMsg(l.PackageName, FuncName, OperationName, operationDescription, nil))
}

// ---------------------------------
func (l LogCollector) RErr(FuncName string, OperationName string, err error, operationDescription string) error {
	//return errMsg + logger.Error
	l.Error(FuncName, OperationName, err, operationDescription)

	return errMsg(FuncName, OperationName, err) //optional return
}

// ---------------------------------
func (l LogCollector) Error(FuncName string, OperationName string, err error, operationDescription string) {

	operationDescription = " fail: " + operationDescription

	loggerToCli.Error(LogMsg(l.PackageName, FuncName, OperationName, operationDescription, err))
	loggerToFile.Error(LogMsg(l.PackageName, FuncName, OperationName, operationDescription, err))

}

// ---------------RWarn------------------
func (l LogCollector) RWarn(FuncName string, OperationName string, err error, operationDescription string) error {
	l.Warn(FuncName, OperationName, err, operationDescription)
	return errMsg(FuncName, OperationName, err)
}

// ----------------RWarnStr-----------------
func (l LogCollector) RWarnStr(FuncName string, OperationName string, errStr string, operationDescription string) error {

	err := errors.New(errStr)

	return l.RWarn(FuncName, OperationName, err, operationDescription)
}

// ----------------Warn-----------------
func (l LogCollector) Warn(FuncName string, OperationName string, err error, operationDescription string) {

	operationDescription = " !!!: " + operationDescription

	loggerToCli.Warn(LogMsg(l.PackageName, FuncName, OperationName, operationDescription, err))
	loggerToFile.Warn(LogMsg(l.PackageName, FuncName, OperationName, operationDescription, err))

}

// -------------------RSlogErr-------------------------
func RSlogErr(packageName, funcName, opName, opDes string, err error) error {
	slog.Error(LogMsg(packageName, funcName, opName, opDes, err))
	return errMsg(funcName, opName, err)
}

//----------------------------------------------------
