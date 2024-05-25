package errorpage

import (
	"html/template"
	"net/http"
)

// =========================================================

func ErrorPage(w http.ResponseWriter, errMsg string) {
	//-------------------------------------------------
	funcName := "ErrorPage"
	//-------------------------------------------------
	templateFile := errorTemplAdrr
	// Define the template
	t := template.Must(template.ParseFiles(templateFile))
	data := struct {
		Msg string
		Url string
	}{
		Msg: errMsg,
		Url: landingPagePath,
	}
	//-----------------------------------
	if err := t.Execute(w, data); err != nil {
		logger.Error(funcName, "Execute-ErrorPage", err, "Execute ErrorPage")
		return
	}
	//-------------------------------------------------
}

// =========================================================
