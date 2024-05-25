package logstack

// ----------------------------------------
func LogMsg(packageName string, FuncName string, OperationName string, description string, RetunedError error) (string, string, string, string, string, string, string, string, string) {
	var errStr string
	errStr = "nil"
	if RetunedError != nil {
		errStr = RetunedError.Error()
	}
	return description, "package", packageName, "func", FuncName, "op", OperationName, "error", errStr

}

//--------------------
