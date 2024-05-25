package handlingreqparameter

import (
	database "main/internal/database"
	query "main/internal/query"
	templatemaker "main/internal/templateMaker"
	errorpage "main/internal/templateMaker/errorpage"
	http "net/http"
	"strconv"
)

// =====================================================
func HandlingReqId(w http.ResponseWriter, r *http.Request, db *database.ArtistsDataBase, callerPkgNames, callerFuncName string) error {
	//------------------Who-am-I?-------------------------------------
	funcName := "HandlingReqId"
	var err error
	//----------------caller-validation------------------------
	if !funcsAccessList[funcName][callerPkgNames][callerFuncName] {
		return errorpage.RErrPageLog(funcName, "logger.FuncCheckAccess", nil, "Check if the caller allowed to use me", w, InternalError500+" : unauthorized function access")
	}
	//-----------What-is-the-reuested-id-Value-----------------
	var artistIdInt int
	artistIdInt, err = whatIsReqIdValueInInt(r)
	if err != nil {
		return errorpage.RErrPageLog(funcName, "WhatIsReqIdValueInInt", err, "retrive id value in int", w, BadRequest400)
	}
	//----------------IsIDExist--------------------------------
	if !query.IsIDExist(db, artistIdInt) {
		return errorpage.RErrPageLog(funcName, "query.IsIDExist", err, "the requested Id does not exist", w, NotFound404)

	}
	//--------------ArtistsPageMakerBasedId---------------------
	var generatedHtml string
	generatedHtml, err = templatemaker.ArtistsPageMakerBasedId(db, artistIdInt)
	if err != nil {
		return logger.RErr(funcName, "templatemaker.ArtistsPageMakerBasedId", err, "creating template")
	}
	//---------write to http.respond----------------------------
	w.Write([]byte(generatedHtml))
	//----------------------------------------------------------
	return nil
}

// =====================================================
func whatIsReqIdValueInInt(r *http.Request) (int, error) {
	funcName := "WhatIsReqIdValueInInt"
	//---------------------------
	IdValueStr := r.URL.Query().Get("id")
	artistIdInt, err := strconv.Atoi(IdValueStr)
	if err != nil {
		return 0, logger.RErr(funcName, "strconv.Atoi", err, "convert string value id to int")
	}
	//--------------------------------
	if artistIdInt < 1 {
		return 0, logger.RErr(funcName, "checkIdvalueValidity", err, "check if ID>1")
	}
	//---------------------------
	return artistIdInt, nil
}

//=====================================================
