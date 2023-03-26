package constants

import "net/http"

var API_META = []struct {
	Code     string
	HttpCode uint
	Message  string
}{
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
	{
		Code:     "E0004",
		HttpCode: http.StatusUnauthorized,
		Message:  "Unknown Secret Key",
	},
	{
		Code:     "E0005",
		HttpCode: http.StatusUnauthorized,
		Message:  "Invalid Secret Key",
	},
	{
		Code:     "E0006",
		HttpCode: http.StatusBadRequest,
		Message:  "Invalid Response Token",
	},

	{
		Code:     "E1000",
		HttpCode: http.StatusBadRequest,
		Message:  "Hostname Already Registered",
	},
	{
		Code:     "E1001",
		HttpCode: http.StatusBadRequest,
		Message:  "Hostname Register Error",
	},
	{
		Code:     "E1002",
		HttpCode: http.StatusBadRequest,
		Message:  "Verified Hostname",
	},
	{
		Code:     "E1003",
		HttpCode: http.StatusBadRequest,
		Message:  "Create Challenge Error",
	},
	{
		Code:     "E1004",
		HttpCode: http.StatusBadRequest,
		Message:  "Challenge Has Been Created",
	},
}
