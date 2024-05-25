package servers

import (
	"net/http"
	"os"
)

func FileServer(staticFilesAddress string) error {
	funcName := "FileServer"
	var opName, opDes string
	var err error
	//-------------------------------
	_, err = os.Stat(staticFilesAddress)
	opName, opDes = "os.State", "directory path  found"
	if os.IsNotExist(err) {
		return logger.RErr(funcName, opName, err, opDes)
	} else {
		logger.Info(funcName, opName, opDes)
	}
	//-------------------------------
	fileServer := http.FileServer(http.Dir(staticFilesAddress))
	http.Handle(cssPathern, http.StripPrefix(cssPathern, fileServer))
	//-------------------------------
	return nil

}
