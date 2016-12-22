package errorhandle

import (
	"net/http"
)

var Err101 = ErrorDefine{http.StatusUnauthorized, 101, "err msg"}
