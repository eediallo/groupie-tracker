package templatemaker

import (
	"bytes"
	"html/template"
	"main/internal/database"
)

// ------------------------------------
func LandingPageMaker(db *database.ArtistsDataBase) error {
	funcName := "LandingPageMaker"
	var opName, opDes string
	var err error
	//--------Parse the template------------------
	var homeTemp *template.Template
	//-------------------------------------------------
	homeTemp, err = template.ParseFiles(hometemplateAddress)
	opName, opDes = "ParseFiles", "parsing Home template"
	if err != nil {
		logger.Error(funcName, opName, err, opDes)
		return logger.ErrMsg(funcName, opName, err)
	} else {
		logger.Info(funcName, opName, opDes)
	}
	//-----------------------------------------------
	var buffer bytes.Buffer // Create a buffer to capture output
	// ----------------------------------------
	data := landingPageMaker_data(db)
	//------------------------------------------
	if err := homeTemp.Execute(&buffer, data); err != nil {
		return logger.RErr(funcName, "homeTemp.Execute", err, "homeTemp-Execute")
	}

	db.TemplatedHomePage = buffer.String()

	return nil
}

// ========================================
func landingPageMaker_data(db *database.ArtistsDataBase) *homePageDataType {
	var artists []artistType
	for _, artist := range db.TableOfArtists {
		artists = append(artists, artistType{
			ID:       artist.ID,
			Name:     artist.Name,
			ImageUrl: artist.ImageUrl,
		})
	}
	//----------------------------------------------
	return &homePageDataType{
		Type:                "artist",
		Artists:             artists,
		CSSHomePath:         homePageCssAdress,
		ArtistIDQueryParam:  artistIDQueryParam,
		QueryParamDelimiter: queryParamDelimiter,
		URLSegmentArtist:    urlSegmentArtist,
		JsFilePath:          JsFilePath,
	}
}

//---------------------------------------
