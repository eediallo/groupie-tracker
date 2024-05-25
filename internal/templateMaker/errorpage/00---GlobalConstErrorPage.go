package errorpage

import (
	logstack "main/internal/logstack"
	//"main"
)

// =================================
const (
	pkgName = "errorpage"

	BadRequest400 = "400 - Bad Request"
	NotFound404   = "404 - Not Found"
	//InternalError500 = "500: server internal error"
	//ArtistNotFound   = "Artist Not Found"
	//--------------------------------
	templateFileAddress = "templates/"
	//homePageFileName    = "index.html"
	//hometemplateAddress = templateFileAddress + homePageFileName
	//artistFileName      = "artist_details.html"
	//artistTempAdrress   = templateFileAddress + artistFileName
	errorFileName  = "error.html"
	errorTemplAdrr = templateFileAddress + errorFileName

	landingPagePath = "/"
	//------------------------------
)

// -------------------------------------
var (
	logger = logstack.LogCollector{
		PackageName: pkgName,
	}
)

// -------------------------------------
