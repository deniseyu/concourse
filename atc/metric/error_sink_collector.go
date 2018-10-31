package metric

import (
	"code.cloudfoundry.org/lager"
)

type ErrorSinkCollector struct {
	logger lager.Logger
}

func NewErrorSinkCollector(logger lager.Logger) ErrorSinkCollector {
	return ErrorSinkCollector{
		logger: logger,
	}
}

func (c *ErrorSinkCollector) Log(f lager.LogFormat) {
	if f.LogLevel != lager.ERROR {
		return
	}

	ErrorLog{
		Value:   1,
		Message: f.Message,
	}.Emit(c.logger)
}
