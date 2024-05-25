package query

import (
	"main/internal/database"
)

// -----------------------------------------------------
func IsIDExist(db *database.ArtistsDataBase, id int) bool {
	//funcName := "IsIDExist"

	for _, artist := range db.TableOfArtists {
		if id == artist.ID {
			return true //, nil
		}
	}

	return false
}

//-----------------------------------------------------
