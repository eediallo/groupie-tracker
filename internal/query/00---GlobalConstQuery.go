package query

import (
	"main/internal/logstack"
)

// --------------------------------
const (
	pkgName = "query"
)

//------------------------------

var (
	logger = logstack.LogCollector{
		PackageName: pkgName,
	}
)

// -------------------------
