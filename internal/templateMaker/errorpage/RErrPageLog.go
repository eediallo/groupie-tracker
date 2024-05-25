package errorpage

import (
	http "net/http"
)

// ======================
func RErrPageLog(funcName string, OperationName string, err error, operationDescription string, w http.ResponseWriter, UserErrMsg string) error {

	ErrorPage(w, UserErrMsg)
	return logger.RErr(funcName, OperationName, err, operationDescription)

}

//-------------------------
