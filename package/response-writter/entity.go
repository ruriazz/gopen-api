package responseWritter

import (
	"github.com/gin-gonic/gin"
	"github.com/ruriazz/gopen-api/package/config"
)

type ResponseWritter struct {
	Config *config.Config
}

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
	Page      int
	Limit     int
	TotalRows int
	TotalPage int
}

type FieldsV1 struct {
	Context  *gin.Context
	MetaCode string
	Message  interface{}
	Data     interface{}
}
