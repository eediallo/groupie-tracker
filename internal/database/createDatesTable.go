package database

//import "fmt"

func (db *ArtistsDataBase) createDatesTable() error {

	funcName := "createDatesTable"

	// -----------------------------------------
	// the structure of Dates.json
	type DatesJson struct {
		Index []struct {
			ID    int      `json:"id"`
			Dates []string `json:"dates"`
		} `json:"index"`
	}
	//---------------------------------------
	var myDatesJsonContent DatesJson
	//---------------------------------------

	if err := fetchUrl(db.TableOfUrls.UrlDates, &myDatesJsonContent); err != nil {
		return logger.RErr(funcName, "fetchUrl", err, "fetching tha Data Url")
	}
	//------initiate map for Table OF data-----------
	db.TableOfDates = make(map[int][]string)
	//------Read data from content and put them in Table-------
	for _, item := range myDatesJsonContent.Index {
		db.TableOfDates[item.ID] = item.Dates
	}
	//---------------------------------------
	return nil
}
