package handlingreqparameter

import (
	"main/internal/database"
	"main/internal/query"

	"main/internal/templateMaker"
	errorpage "main/internal/templateMaker/errorpage"
	"net/http"
)

// =========================================
func HandlingReqName(w http.ResponseWriter, r *http.Request, db *database.ArtistsDataBase) error {
	funcName := "HandlingReqName"
	//var opName, opDes string
	var err error
	// -----------What-is-the-reuested-name-Value-----
	nameValueStr := r.URL.Query().Get("name")

	// ---------------------------------------------
	var artistIdInt int
	artistIdInt, err = query.IdFromName(db, nameValueStr)
	if err != nil {

		return errorpage.RErrPageLog(funcName, "query.IdBasedName", nil, "retrive id from Name", w, NotFound404+" : artist id does not exist")

	}
	//----------------------------
	var generatedHtml string
	generatedHtml, err = templatemaker.ArtistsPageMakerBasedId(db, artistIdInt)
	if err != nil {
		return logger.RErr(funcName, "templatemaker.ArtistsPageMakerBasedId", err, "creating template")
	}

	w.Write([]byte(generatedHtml))
	//----------------------------
	return nil
}

//=========================================
