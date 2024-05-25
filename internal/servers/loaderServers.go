package servers

import (
	"main/internal/database"
	// "main/internal/logStack"
)

func LoadServerS(db *database.ArtistsDataBase) error {
	//---------------------------------------

	funcName := "LoadServerS"
	var opName, opDes string
	//---------------------Front-End-------------------------
	opName, opDes = "FileServer", "fetching data "
	if err := FileServer(staticFilesAddress); err != nil {
		return logger.RErr(funcName, opName, err, opDes)
	} else {
		logger.Info(funcName, opName, opDes)
	}
	//-------------------------------
	opName, opDes = "LoadingWebServer", "Loading web server"
	if err := LoadingWebServer(db); err != nil {
		return logger.RErr(funcName, opName, err, opDes)
	} else {
		logger.Info(funcName, opName, opDes)
	}
	//-------------------------------
	return nil
}
