package landinghandlers

import (
	"net/http"
)

// ==============================================
func hasUrljustOneParameter(r *http.Request) bool {
	//funcName := "hasItjustOneParameter"

	//----------------------
	var paramCount int
	query := r.URL.Query()
	for k := range query {
		_ = k
		paramCount++
	}

	return paramCount == 1
}

// ========================================================
func requestedParameter(r *http.Request) (string, error) {
	funcName := "hasItjustOneParameter"
	//--------------------------
	//var paramCount int
	query := r.URL.Query()
	for key := range query {

		return key, nil
	}
	return "", logger.RWarn(funcName, "requestedParameter", nil, "requestedParameter")
}

// ===================================
