package query

import (
	"main/internal/database"
	"strings"
)

// ==============================================

func IdFromName(db *database.ArtistsDataBase, name string) (int, error) {
	funcName := "IdBasedName"
	var err error
	// -------------------------------------
	var isIdFound bool //default value false
	var id int
	//---------------------------------------
	for _, item := range db.TableOfArtists {
		if strings.ToLower(name) == strings.ToLower(item.Name) {

			isIdFound = true
			id = item.ID
		}
	}
	//---------------------------------
	if !isIdFound {

		return 0, logger.RErr(funcName, "artist.ID == 0", err, "check if its empty")
	}
	//---------------------------------
	return id, nil
}

// ==============================================
