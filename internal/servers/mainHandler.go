package servers

import (
	"main/internal/handlers/landinghandlers"
	"net/http"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	funcName := "MainHandler"
	var opName, opDes string
	//-------------------------------
	switch {
	case r.URL.Path == landingPagePath && r.Method == http.MethodGet: // "/"

		opName, opDes = "HomeHandler", "found Home page handler template"
		if err := landinghandlers.HomeHandler(w, r); err != nil {
			logger.Error(funcName, opName, err, opDes)
			return
		} else {
			logger.Info(funcName, opName, opDes)
		}
		//-------------------------------
	case r.URL.Path == artistPageAdrress && r.Method == http.MethodGet: // "/static/"
		logger.Info(funcName, "ArtistHandler", "Artist template found")
		if err := landinghandlers.ArtistHandler(w, r); err != nil {
			logger.Warn(funcName, "ArtistHandler", err, "Artist template found")

		}
		//-------------------------------
	default:
		logger.Info(funcName, "default", "user try to GET unknown page")
		//-------------------------------
	}
	//-------------------------------
}

//-----------------
