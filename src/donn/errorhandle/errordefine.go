package errorhandle

import ()

type ErrorDefine struct {
	httpcode int
	Code     int    `form:"code" json:"code" binding:"required"`
	Message  string `form:"msg" json:"msg" binding:"required"`
}

//func NewErrorDefine(code int, message string) *ErrorDefine {
//	ed := new(ErrorDefine)
//	ed.code = code
//	ed.message = message
//	return ed
//}
