package database

import (
	"main/internal/logstack"
)

// ----------------------------------------
const (
	pkgName    = "DataBase"
	mainApiUrl = `https://groupietrackers.herokuapp.com/api`
)

// -----------------------------------
var (
	//myUrlTable UrlTable
	//-------logger instance---------------
	logger = logstack.LogCollector{
		PackageName: pkgName,
	}
)

// ------------------------------------------------
// the structure of mainApiUrl.json
type UrlTable struct {
	UrlArtists   string `json:"artists"`
	UrlLocations string `json:"locations"`
	UrlDates     string `json:"dates"`
	UrlRelation  string `json:"relation"`
}

// --------------------------------
// define the structure of Artist type for TableOfArtists in database
type Artist struct {
	ID                 int                 `json:"id"`
	Name               string              `json:"name"`
	ImageUrl           string              `json:"image_url"`
	CreationDate       int                 `json:"creation_date"`
	FirstAlbum         string              `json:"first_album"`
	Members            []string            `json:"members"`
	Locations          []string            `json:"locations"`
	Dates              []string            `json:"dates"`
	Relation           map[string][]string `json:"relation"`
	ArtistIdQueryParam string              `json:"artistIdQueryParam"`
}

// ----------------------------------
// the structure of a database
type ArtistsDataBase struct {
	TableOfArtists    map[int]Artist
	TableOfMembers    map[int][]string
	TableOfLocations  map[int][]string
	TableOfDates      map[int][]string
	TableOfRelation   map[int]map[string][]string
	TableOfUrls       UrlTable
	TemplatedHomePage string
}

// ----------------------------------------
