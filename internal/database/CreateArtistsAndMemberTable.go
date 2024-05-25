package database

//import "fmt"

// -------------------------------------------
func (db *ArtistsDataBase) CreateArtistsAndMemberTable() error {

	funcName := "CreateArtistsAndMemberTable"
	// ---------------------------------------
	// the structure as helper to read json data
	type ArtistsJson []struct {
		ID           int      `json:"id"`
		Image        string   `json:"image"`
		Name         string   `json:"name"`
		Members      []string `json:"members"`
		CreationDate int      `json:"creationDate"`
		FirstAlbum   string   `json:"firstAlbum"`
		Locations    string   `json:"locations"`
		ConcertDates string   `json:"concertDates"`
		Relations    string   `json:"relations"`
	}
	//---------------------------------------------
	var myArtistsJsonContent ArtistsJson
	//-----------------------------

	if err := fetchUrl(db.TableOfUrls.UrlArtists, &myArtistsJsonContent); err != nil {
		return logger.RErr(funcName, "fetchUrl", err, "fetching UrlArtists")
	}

	//---------initiate empty map for Tables------------------
	db.TableOfArtists = make(map[int]Artist)
	db.TableOfMembers = make(map[int][]string)
	//------Read data from content and put them in Table-------
	for _, item := range myArtistsJsonContent {

		myArtistData := Artist{
			ID:           item.ID,
			Name:         item.Name,
			ImageUrl:     item.Image,
			CreationDate: item.CreationDate,
			FirstAlbum:   item.FirstAlbum,
		}

		db.TableOfArtists[item.ID] = myArtistData
		db.TableOfMembers[item.ID] = item.Members
	}
	//----------------------------------
	return nil
}

//--------------------------------
