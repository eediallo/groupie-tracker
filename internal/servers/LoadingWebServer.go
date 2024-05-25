package servers

import (
	"main/internal/database"
	"main/internal/handlers/landinghandlers"
	"main/internal/templateMaker"
	"net/http"
)

func LoadingWebServer(db *database.ArtistsDataBase) error {
	funcName := "LoadingWebServer"

	//-------------------------------
	serverWebGuiApp := http.Server{
		Addr:    portAdrressNumber,
		Handler: nil,
	}
	landinghandlers.SetDbAddress(db)
	templatemaker.LandingPageMaker(db)

	http.HandleFunc(landingPagePath, MainHandler)

	if err := serverWebGuiApp.ListenAndServe(); err != nil {
		return logger.RErr(funcName, "ListenAndServe", err, "Error in running web-server")
	}
	//-------------------------------
	return nil
}
