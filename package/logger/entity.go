package logger

type ResponseLog struct {
	Hostname       string                 `json:"hostname"`
	Fullpath       string                 `json:"fullpath"`
	Method         string                 `json:"method"`
	Parameters     map[string]interface{} `json:"parameters"`
	Body           map[string]interface{} `json:"body"`
	Headers        map[string][]string    `json:"headers"`
	UserAgent      string                 `json:"user_agent"`
	Latency        string                 `json:"latency"`
	Response       interface{}            `json:"response"`
	ResponseStatus int                    `json:"response_status"`
}
