package errors

var enErrorMsgMap = map[ErrorCode]string{
	CodeInternalError:     "system error",
	CodeInvalidParamError: "invalid param",
}

func GetEnErrorMsg(ec ErrorCode) string {
	return enErrorMsgMap[ec]
}
