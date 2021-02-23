package model

const (
	LogLevelError       = "error"
	LogLevelWarning     = "warn"
	LogLevelInformation = "info"
	LogLevelDebug       = "debug"

	LogTypeMiddleware = "mid"
	LogTypeAPI        = "api"
	LogTypeRouter     = "router"
	LogTypeService    = "serv"
	LogTypeRepository = "repo"

	LogStatusSuccess = "success"
	LogStatusFailed  = "failed"
	LogStatusData    = "data"
	LogStatusRequest = "request"
)

type LogBody struct {
	Method string `json:"method"`
	Path   string `json:"path"`
	Error  string `json:"error_message,omitempty"`
}
