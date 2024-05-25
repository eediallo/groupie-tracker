package landinghandlers

import (
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) error {
	//-----------------------------------------------------------
	funcName := "HomeHandler"
	// //-------------------------------------------

	logger.Info(funcName, "call-HomeHandler", "User opens Home Page")
	//-----------------------------------------------------------

	//----------------------------------
	w.Write([]byte(db.TemplatedHomePage))
	//---------------------------------------------------------
	return nil
}
