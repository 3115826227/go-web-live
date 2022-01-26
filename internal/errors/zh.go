package errors

var zhErrorMsgMap = map[ErrorCode]string{
	CodeInternalError:     "系统错误",
	CodeInvalidParamError: "参数错误",
}

func GetZhErrorMsg(ec ErrorCode) string {
	return zhErrorMsgMap[ec]
}
