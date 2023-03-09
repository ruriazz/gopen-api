package responseHelper

import (
	"github.com/gin-gonic/gin"
)

type BaseResponseV1 struct {
	Meta       MetaV1        `json:"meta,omitempty"`
	Message    interface{}   `json:"message,omitempty"`
	Data       interface{}   `json:"data,omitempty"`
	Pagination *PaginationV1 `json:"pagination,omitempty"`
}

type MetaV1 struct {
	Success bool   `json:"success"`
	Latency string `json:"latency"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

type PaginationV1 struct {
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	TotalRow  int `json:"totalRow"`
	TotalPage int `json:"totalPage"`
}

type FieldsV1 struct {
	Context  *gin.Context
	MetaCode string
	Message  interface{}
	Data     interface{}
}
