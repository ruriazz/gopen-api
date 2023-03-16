package settings

import "net/http"

var API_META = []ApiMeta{
	{
		Code:     "S0000",
		HttpCode: http.StatusOK,
		Message:  "OK",
	},
	{
		Code:     "S0001",
		HttpCode: http.StatusOK,
		Message:  "Data Not Found",
	},

	{
		Code:     "E0000",
		HttpCode: http.StatusBadRequest,
		Message:  "General Error",
	},
	{
		Code:     "E0001",
		HttpCode: http.StatusBadRequest,
		Message:  "Query Parameter Validation Error",
	},
	{
		Code:     "E0002",
		HttpCode: http.StatusBadRequest,
		Message:  "Request Body Validation Error",
	},
	{
		Code:     "E0003",
		HttpCode: http.StatusBadRequest,
		Message:  "Fetch Data Error",
	},
}
