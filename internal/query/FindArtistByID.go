package query

import (
	"main/internal/database"
)

// -----------------------------------------------------
func FindArtistByID(db *database.ArtistsDataBase, id int) (database.Artist, error) {
	funcName := "FindArtistByID"
	var err error
	//----------------------------------
	var artist database.Artist
	//------------------------------------
	var isIdFound bool //default value false

	for _, item := range db.TableOfArtists { //item=artist
		if id == item.ID {

			isIdFound = true //flag

			artist.ID = item.ID
			artist.Name = item.Name
			artist.CreationDate = item.CreationDate
			artist.FirstAlbum = item.FirstAlbum
			artist.Relation = db.TableOfRelation[artist.ID]
			artist.Members = db.TableOfMembers[artist.ID]
			artist.ImageUrl = item.ImageUrl

		}
	}
	//------------------check if its empty-------------
	if !isIdFound {
		return artist, logger.RErr(funcName, "artist.ID == 0", err, "check if its empty")
	}
	//-------------------------------------
	return artist, nil
}

//=======================================================
