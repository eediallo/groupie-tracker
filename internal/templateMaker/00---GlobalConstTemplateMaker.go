package templatemaker

import (
	//"main/internal/database"
	"main/internal/logstack"
)

const (
	pkgName = "templatemaker"

	rootDir             = "../.."
	artistIDQueryParam  = "id"
	queryParamDelimiter = "?"
	urlSegmentArtist    = "/artist"
	statisDirName       = "/static"
	CssDirName          = "/styles"
	artistCSSFileName   = "/artist.css"
	homeCSSFileName     = "/index.css"
	landingPagePath     = "/"

	templateFileAddress = "templates/"
	homePageFileName    = "index.html"

	JsFilePath = "/static/js/searchName.js"

	artistFileName = "artist_details.html"

	//----------------------------------------------
	hometemplateAddress  = templateFileAddress + homePageFileName
	staticsDirAdd        = rootDir + statisDirName     // --> "../../static"
	CssDirAdd            = staticsDirAdd + CssDirName  // --> "../../static/styles"
	homePageCssAdress    = CssDirAdd + homeCSSFileName // -->"../../static/styles/index.css"
	artistCSSPageAddress = CssDirAdd + artistCSSFileName

	artistTempAdrress = templateFileAddress + artistFileName
)

// -----------------------------------
var (
	//-------logger instance---------------
	logger = logstack.LogCollector{
		PackageName: pkgName,
	}
)

// -----------------------------
// ---------------------------------
type artistType struct {
	ID       int
	Name     string
	ImageUrl string
}

// -------------------------------------------
type homePageDataType struct { //landingpage
	Type                string
	Artists             []artistType
	CSSHomePath         string
	ArtistIDQueryParam  string
	QueryParamDelimiter string
	URLSegmentArtist    string
	JsFilePath          string
}

// -------------------------------------
type artistPageByIdDataType struct { //artist page
	Type string
	//Artist             database.Artist // []database.Artist
	Id                   int
	Name                 string
	CreationDate         int
	FirstAlbum           string
	Relation             map[string][]string
	Members              []string
	CSSPageAddress       string
	ImageUrl             string
	HomePage             string // "/"
	ArtistCSSPageAddress string
}

//----------------------------------------
