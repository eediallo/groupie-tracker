package logstack

import "errors"

// ----------------------------------------
func errMsg(FuncName string, Operation string, RetunedError error) error {
	var RetunedErrorStr string
	if RetunedError != nil {
		RetunedErrorStr = RetunedError.Error()
	}

	errStr := "<==={" + pkgName + "}--" + FuncName + "<---" + Operation + "[" + RetunedErrorStr + "]***"

	return errors.New(errStr)
}

// ----------------------------------------
