package database

// ======================================================
func (db *ArtistsDataBase) createRelationsTable() error {
	funcName := "createDatesTable"
	//var opName, opDes string

  	//var err error
	//--------------------setup------------------------------
	// the structure of relation.json
	type relationsJson struct {
		Index []struct {
			ID             int `json:"id"`
			DatesLocations map[string][]string
		} `json:"index"`
	}
	//-----------------------------------------
	var myRelationsJsonContent relationsJson
	//-------------------------------------------------

  	if err := fetchUrl(db.TableOfUrls.UrlRelation, &myRelationsJsonContent); err != nil {

		return logger.RErr(funcName, "fetchUrl", err, "fetching Url Relation")
	}

	//--------------------------------------
	//initiate the map (create an empty map)
	db.TableOfRelation = make(map[int]map[string][]string)
	//------Read data from content and put them in Table-------
	for _, item := range myRelationsJsonContent.Index {
		db.TableOfRelation[item.ID] = item.DatesLocations
	}
	//----------------------------------
	return nil
}

//======================================================
