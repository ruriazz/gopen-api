package logger

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ruriazz/gopen-api/src/constants"
	log "github.com/sirupsen/logrus"
)

type executionLogAttributeV1 struct {
	ModuleType    string `json:"module_type"`
	ModulePackage string `json:"module_package"`
	ModuleName    string `json:"module_name"`
	Method        string `json:"method"`
	ExecutableID  int    `json:"executable_id"`
	Message       string `json:"message"`
}

type contextDataAttribute struct {
	Hostname   string                 `json:"hostname"`
	Fullpath   string                 `json:"fullpath"`
	Method     string                 `json:"method"`
	Parameters map[string]interface{} `json:"parameters"`
	Body       map[string]interface{} `json:"body"`
	Headers    map[string][]string    `json:"headers"`
	UserAgent  string                 `json:"user_agent"`
}

type ExecutionLog interface {
	Info(method string, executableID int, message string, context *gin.Context)
	Warning(method string, executableID int, message string, context *gin.Context)
	Error(method string, executableID int, message string, context *gin.Context)

	getLogEntry(method string, executableID int, message string, context *gin.Context) *log.Entry
	parseContext(context *gin.Context) contextDataAttribute
}

type executionLog struct {
	ModuleType    constants.LogModuleType
	ModulePackage string
	ModuleName    string
}

func NewExecutionLog(moduleType constants.LogModuleType, modulePackage string, moduleName string) ExecutionLog {
	return executionLog{
		ModuleType:    moduleType,
		ModulePackage: modulePackage,
		ModuleName:    moduleName,
	}
}

func (l executionLog) Info(method string, executableID int, message string, context *gin.Context) {
	logEntry := l.getLogEntry(method, executableID, message, context)
	logEntry.Infof("%s_call_trace", l.ModuleType)
}

func (l executionLog) Warning(method string, executableID int, message string, context *gin.Context) {
	logEntry := l.getLogEntry(method, executableID, message, context)
	logEntry.Warningf("%s_call_trace", l.ModuleType)
}

func (l executionLog) Error(method string, executableID int, message string, context *gin.Context) {
	logEntry := l.getLogEntry(method, executableID, message, context)
	logEntry.Errorf("%s_call_trace", l.ModuleType)
}

func (l executionLog) getLogEntry(method string, executableID int, message string, context *gin.Context) *log.Entry {
	var contextData contextDataAttribute
	if context != nil {
		contextData = l.parseContext(context)
	}

	return log.WithFields(log.Fields{
		"attributes": &executionLogAttributeV1{
			ModuleType:    l.ModuleType.String(),
			ModulePackage: l.ModulePackage,
			ModuleName:    l.ModuleName,
			Method:        method,
			ExecutableID:  executableID,
			Message:       message,
		},
		"context_data": contextData,
	})
}

func (l executionLog) parseContext(context *gin.Context) contextDataAttribute {
	url, _ := url.Parse(fmt.Sprintf("http://%s", context.Request.Host))
	param := make(map[string]interface{}, len(context.Request.URL.Query()))
	for k, v := range context.Request.URL.Query() {
		param[k] = v[0]
	}
	if len(param) == 0 {
		param = nil
	}

	return contextDataAttribute{
		Hostname:   strings.TrimPrefix(url.Hostname(), "www."),
		Fullpath:   context.FullPath(),
		Method:     context.Request.Method,
		Parameters: param,
		Headers:    context.Request.Header,
		UserAgent:  context.Request.UserAgent(),
	}
}
