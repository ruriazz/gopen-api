package responseWritter

import (
	"fmt"
	"time"

	"github.com/ruriazz/gopen-api/package/config"
)

func CreateResponseWritter(config *config.Config) (*ResponseWritter, error) {
	return &ResponseWritter{
		Config: config,
	}, nil
}

func (rw *ResponseWritter) JsonResponse(fields FieldsV1) {
	response := &BaseResponseV1{}

	if fields.MetaCode == "" {
		fields.MetaCode = "S0000"
	}

	metaObject := rw.Config.ApiMetaByCode(fields.MetaCode)
	if metaObject == nil {
		panic(fmt.Errorf("invalid API Meta with code '%s'", fields.MetaCode))
	}

	rt := fields.Context.GetTime("requestTime")

	response.Meta = MetaV1{
		Success: metaObject.Http >= 200 && metaObject.Http <= 299,
		Code:    metaObject.Code,
		Message: metaObject.Message,
		Latency: (time.Since(rt) * 1000).String(),
	}

	if fields.Data != nil {
		response.Data = fields.Data
	} else if fields.Message != nil {
		response.Message = fields.Message
	}

	fields.Context.JSON(metaObject.Http, response)
}
