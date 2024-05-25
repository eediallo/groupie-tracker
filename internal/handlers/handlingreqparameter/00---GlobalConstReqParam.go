package handlingreqparameter

import (
	"main/internal/logstack"
)

// =================================
// ----------------------------------------

// -----------------const-----------------------
const (
	pkgName = "handlingreqparameter"

	BadRequest400    = "400 - Bad Request"
	NotFound404      = "404 - Not Found"
	InternalError500 = "500: server internal error"
)

// -----------variable--------------------
var (
	logger = logstack.LogCollector{
		PackageName: pkgName,
	}
	//-------------------------
	funcsAccessList = map[string]map[string]map[string]bool{
		"HandlingReqId": { //it's me
			"landinghandlers": { //packageName
				"ArtistHandler": true,  //this func can use me
				"func2_test":    false, //this func can not use me
			},
		},
	}
	//-------------------------
)

//===================================
