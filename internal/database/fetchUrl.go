package database

import (
	"encoding/json"
	//"errors"
	"io"
	"net/http"
)

// ----------------------------------------------
func fetchUrl(url string, structure interface{}) error {
	funcName := "fetchUrl"

	//--------GET API CALL----------------
	resp, err := http.Get(url)

	if err != nil {
		return logger.RErr(funcName, "http.Get", err, "http.Get to Url")
	}
	defer resp.Body.Close()
	//-------------------------------------
	// Check response status
	if resp.StatusCode != http.StatusOK {

		return logger.RWarnStr(funcName, "resp.StatusCode", "resp.StatusCode is not OK=200", "Check response status")
	}
	//----------------read response body--------------------
	contentByte, err := io.ReadAll(resp.Body)
	if err != nil {
		return logger.RErr(funcName, "io.ReadAll", err, "reading body of http res")
	}
	//--------------Unmarshal----------

	if json.Unmarshal(contentByte, &structure) != nil {
		return logger.RErr(funcName, "json.Unmarshal", err, "Unmarshalling Json file")
	}
	//--------------------------
	return nil
}
