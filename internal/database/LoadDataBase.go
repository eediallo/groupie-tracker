package database

//import "fmt"

func LoadDataBase(db *ArtistsDataBase, rootDir string) error {
	funcName := "LoadDataBase"
	//-----------------------------------
	if err := fetchUrl(mainApiUrl, &db.TableOfUrls); err != nil {
		return logger.RErr(funcName, "fetchUrl", err, "fetching tha main Url")

	}
	//------------CreateArtistsAndMemberTable---------------

	if err := db.CreateArtistsAndMemberTable(); err != nil {
		return logger.RErr(funcName, "CreateArtistsAndMemberTable", err, "Creating Artists And Member Table")
	}
	//--------------createLocationTable-------------

	if err := db.createLocationTable(); err != nil {
		return logger.RErr(funcName, "createLocationTable", err, "creating Location Table")
	}
	//-----------------createDatesTable-------------

	if err := db.createDatesTable(); err != nil {
		return logger.RErr(funcName, "createDatesTable", err, "creating Dates Table")
	}
	//---------createRelationsTable---------------

	if err := db.createRelationsTable(); err != nil {
		return logger.RErr(funcName, "createRelationsTable", err, "creating Relations Table")
	}

	//------------------------------------------
	return nil
}
