package landinghandlers

import (
	//"errors"
	handlingreqparameter "main/internal/handlers/handlingreqparameter"
	errorpage "main/internal/templateMaker/errorpage"
	http "net/http"
)

// ===================================
func ArtistHandler(w http.ResponseWriter, r *http.Request) error {
	funcName := "ArtistHandler"
	var err error
	//------------------------------------
	logger.Info(funcName, "ArtistHandler", "client requests artist details")
	//--------------has it just 1 parameter-----------
	if !hasUrljustOneParameter(r) {

		return errorpage.RErrPageLog(funcName, "hasUrljustOneParameter", nil, "has Url just One Parameter", w, badRequest400)
	}
	//---------------requested-Parameter---------------------------
	var requestedParam string
	requestedParam, err = requestedParameter(r)
	if err != nil {

		return errorpage.RErrPageLog(funcName, "requestedParameter", nil, "what is the requested parameter", w, badRequest400)

	}
	//-------------------requested Parameter Switch -------------------
	switch requestedParam {
	case "id":
		if handlingreqparameter.HandlingReqId(w, r, db, pkgName, funcName) != nil {
			return logger.RWarn(funcName, "HandlingReqId", err, "to create and write template based on requested ID")
		}
		//---------------------
	case "name":

		handlingreqparameter.HandlingReqName(w, r, db)

	default:

		return errorpage.RErrPageLog(funcName, "default-Switch-case", nil, "default-Switch-case", w, badRequest400)

	}

	//---------------------------------------------
	return nil
}

//======================================================
