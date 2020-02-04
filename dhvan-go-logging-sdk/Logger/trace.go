package Logger

import (
	"dhvan-go-logging-sdk/enums"
)

type trace struct {
	next  abstractLogger
	level enums.LogLevel
}

func (trc *trace) setNext(next abstractLogger) {
	trc.next = next
}

func (trc *trace) setLevel(level enums.LogLevel) {
	trc.level = level
}

func (trc *trace) Execute(fluentdLogger *FluentdLogger, passedLogLevel enums.LogLevel, tag string, data map[string]string) error {
	var fluentdPostError error
	if trc.level == passedLogLevel && passedLogLevel >= enums.GetLogLevelFromLogType(fluentdLogger.InitLogDetails.GlobalLoggingType) {
		fluentdPostError = fluentdLogger.FluentdConnection.Post(tag, data)
	}
	if trc.next != nil {
		chainErr := trc.next.Execute(fluentdLogger, passedLogLevel, tag, data)
		if chainErr != nil {
			return chainErr
		}
	}

	return fluentdPostError
}
