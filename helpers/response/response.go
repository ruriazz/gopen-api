package responseHelper

import (
	"fmt"
	"time"

	"github.com/ruriazz/gopen-api/package/settings"
)

func JSON(fields FieldsV1) {
	response := &BaseResponseV1{}

	if fields.MetaCode == "" {
		fields.MetaCode = "S0000"
	}

	metaObject := apiMetaByCode(fields.MetaCode)
	if metaObject == nil {
		panic(fmt.Errorf("invalid API Meta with code '%s'", fields.MetaCode))
	}

	rt := fields.Context.GetTime("requestTime")

	response.Meta = MetaV1{
		Success: metaObject.HttpCode >= 200 && metaObject.HttpCode <= 299,
		Code:    metaObject.Code,
		Message: metaObject.Message,
		Latency: time.Since(rt).String(),
	}

	if fields.Data != nil {
		response.Data = fields.Data
	} else if fields.Message != nil {
		response.Message = fields.Message
	}

	if fields.Pagination != nil {
		response.Pagination = fields.Pagination
	}

	fields.Context.JSON(int(metaObject.HttpCode), response)
	fields.Context.Next()
}

func apiMetaByCode(code string) *settings.ApiMeta {
	for _, meta := range settings.API_META {
		if meta.Code == code {
			return &meta
		}
	}

	return nil
}
