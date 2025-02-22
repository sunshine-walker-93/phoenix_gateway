package constant

type ErrorStruct struct {
	Code int32
	Msg  string
}

// 错误码的前四位是模块编号，5-8位是模块内的错误编号（5-6表示错误类型，7-8表示具体错误）

var Success = ErrorStruct{0, "成功"}

var Error = ErrorStruct{10010100, "未知错误"}
var InvalidParams = ErrorStruct{10010101, "参数非法"}
var ErrorRegisterFailed = ErrorStruct{10010102, "注册失败"}

var ErrorAuthCheckTokenFail = ErrorStruct{10010200, "Token鉴权失败"}
var ErrorAuthCheckTokenTimeout = ErrorStruct{10010201, "Token鉴权失败"}
var ErrorAuthToken = ErrorStruct{10010202, "Token鉴权失败"}

var ErrorUploadCheckImageFormat = ErrorStruct{10010300, "上传的图片大小不合法"}
var ErrorUploadCheckImageExt = ErrorStruct{10010300, "上传的图片格式不合法"}
