package servers

import (
	"main/internal/logstack"
)

// ---------------------------------------------------
const (
	pkgName           = "servers"
	landingPagePath   = "/"
	artistPagePath    = "artist"
	artistPageAdrress = landingPagePath + artistPagePath

	//badRequest400    = "400 - Bad Request"
	//NotFound404      = "404 - Not Found"
	//InternalError500 = "500: server internal error"

	cssPathern = "/static/"
	//mainFolder = "./"
	LogFilesDirectory = "logs/"

	staticFilesAddress = "static"
	portNumber         = "8080"
	portAdrressNumber  = ":" + portNumber
)

// ------------------loggere-instance-----------
var (
	logger = logstack.LogCollector{
		PackageName: pkgName,
	}

	//templatedHomePage string
)

//-------------------------------------------
