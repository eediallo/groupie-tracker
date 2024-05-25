package database

//import "fmt"

func (db *ArtistsDataBase) createLocationTable() error {
	funcName := "createDatesTable"
	//var opName, opDes string
	//var err error
	//--------------------------------

	// the structure of locations.json
	type LocationsJson struct {
		Index []struct {
			ID        int      `json:"id"`
			Locations []string `json:"locations"`
			Dates     string   `json:"dates"`
		} `json:"index"`
	}

	//--------------------------------

	var myLocationJsonContent LocationsJson
	//--------------------------------

	if err := fetchUrl(db.TableOfUrls.UrlLocations, &myLocationJsonContent); err != nil {
		return logger.RErr(funcName, "fetchUrl", err, "fetching Url Locations")
	}
	//---------------------------------
	db.TableOfLocations = make(map[int][]string)
	//------Read data from content and put them in Table-------
	for _, item := range myLocationJsonContent.Index {
		db.TableOfLocations[item.ID] = item.Locations
	}
	//-------------------------------
	return nil
}
