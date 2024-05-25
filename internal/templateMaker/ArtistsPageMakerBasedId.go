package templatemaker

import (
	"bytes"
	//"fmt"
	"html/template"
	"main/internal/database"
	"main/internal/query"
)

// =======================================================
func ArtistsPageMakerBasedId(db *database.ArtistsDataBase, artistIDInt int) (string, error) {
	funcName := "ArtistsPageMakerBasedId"
	var opName, opDes string
	var err error
	//------------------------------
	var artistData *artistPageByIdDataType
	artistData, err = artistPageMakerBasedId_data(db, artistIDInt)
	if err != nil {
		return "", logger.RErr(funcName, "artistPageMakerBasedId_data", err, "retrive data from db based of ID")
	}
	//---------------------
	//Parse the template
	var ArtistTemp *template.Template
	ArtistTemp, err = template.ParseFiles(artistTempAdrress)
	opName, opDes = "ParsingArtistTemp", "parsing Artist Template"
	if err != nil {
		return "", logger.RErr(funcName, opName, err, opDes)
	} else {
		logger.Info(funcName, opName, opDes)
	}
	//-----------------------------------------------
	var buffer bytes.Buffer // Create a buffer to capture output

	//------------------------------------------
	if ArtistTemp.Execute(&buffer, artistData) != nil {
		return "", logger.RErr(funcName, "ArtistTemp.Execute", err, "Executing artistData(based Id) to buffer")
	}
	//----------------------------
	return buffer.String(), nil
}

// =======================================================
func artistPageMakerBasedId_data(db *database.ArtistsDataBase, artistIDInt int) (*artistPageByIdDataType, error) {
	//-----------------------------------------------
	funcName := "artistPageMaker_data"
	var opName, opDes string
	var err error
	//----------------------------------------------
	// Retrieve artist details from the database based on the name
	var artist database.Artist
	artist, err = query.FindArtistByID(db, artistIDInt)
	opName, opDes = "FindArtistByID", "findind artists by its ID"
	if err != nil {
		return nil, logger.RErr(funcName, opName, err, opDes)
	} else {
		logger.Info(funcName, opName, opDes)
	}
	//--------------------------------------------
	_ = artist
	data := &artistPageByIdDataType{
		Type:                 "artist",
		Id:                   artist.ID,
		Name:                 artist.Name,
		CreationDate:         artist.CreationDate,
		FirstAlbum:           artist.FirstAlbum,
		Relation:             artist.Relation,
		Members:              artist.Members,
		ImageUrl:             artist.ImageUrl,
		ArtistCSSPageAddress: artistCSSPageAddress,
		HomePage:             landingPagePath,
	}
	//
	//------------------------------
	return data, nil
}

//=======================================================
