package settings

import "net/http"

var API_META = []ApiMeta{
	{
		Code:     "S0000",
		HttpCode: http.StatusOK,
		Message:  "OK",
	},
	{
		Code:     "E0000",
		HttpCode: http.StatusBadRequest,
	},
}
