package responseHelper

import (
	"github.com/gin-gonic/gin"
	paginationHelper "github.com/ruriazz/gopen-api/src/helpers/pagination"
)

type BaseResponseV1 struct {
	Meta       MetaV1                         `json:"meta,omitempty"`
	Message    interface{}                    `json:"message,omitempty"`
	Data       interface{}                    `json:"data,omitempty"`
	Pagination *paginationHelper.PaginationV1 `json:"pagination,omitempty"`
}

type MetaV1 struct {
	Success bool   `json:"success"`
	Latency string `json:"latency"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

type FieldsV1 struct {
	Context    *gin.Context
	MetaCode   string
	Message    interface{}
	Data       interface{}
	Pagination *paginationHelper.PaginationV1
	Error      error
}
