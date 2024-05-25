package main

import (
	"log/slog"
	"main/internal/database"
	//"main/internal/handlers"
	"main/internal/logstack"
	"main/internal/servers"
	//"os"
) //

const (
	pkgName      = "main"
	rootFromMain = "../.."
)

// ------------------logger-instance------------
var (
	logger = logstack.LogCollector{
		PackageName: pkgName,
	}
)
//--------------------

func main() {
	funcName := "main"
	var opName, opDes string
	//var err error
	//fmt.Println(os.Getwd())
	//------------------
	//err = logstack.LoadLoggerStack()
	opName, opDes = "LoadLoggerStack", "load log stack"
	if err := logstack.LoadLoggerStack(); err != nil {
		slog.Error(logstack.LogMsg(logger.PackageName, funcName, opName, opDes, err))
		return
	} else {
		slog.Info(logstack.LogMsg(logger.PackageName, funcName, opName, opDes, err))
	}
	//-------------createDB-------------
	MyDataBase := database.CreateDB()
	//----------------------------
	//err = database.LoadDataBase(&MyDataBase, rootFromMain)
	opName, opDes = "LoadDataBase", "loading Database"
	if err := database.LoadDataBase(&MyDataBase, rootFromMain); err != nil {
		logger.Error(funcName, opName, err, opDes)
		return
	} else {
		logger.Info(funcName, opName, opDes)
	}
	//------------------------
	opName, opDes = "LoadServerS", "loading servers"
	if err := servers.LoadServerS(&MyDataBase); err != nil {
		logger.Error(funcName, opName, err, opDes)
		return
	} else {
		logger.Info(funcName, opName, opDes)
	}
	//------------------------
}
