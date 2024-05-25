package landinghandlers

import (
	"main/internal/database"
	"main/internal/logstack"
	//"main"
)

// var db =
const (
	pkgName         = "landinghandlers"
	landingPagePath = "/"

	badRequest400 = "400 - Bad Request"

	myDomain          = "http://localhost"
	portNumber        = "8080"
	portAdrressNumber = ":" + portNumber
	Url               = myDomain + portAdrressNumber
)

// ------------------loggere------------
var (
	logger = logstack.LogCollector{
		PackageName: pkgName,
	}

	db *database.ArtistsDataBase
)

//---------------------------
