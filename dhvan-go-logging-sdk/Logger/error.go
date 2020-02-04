package Logger

import (
	"dhvan-go-logging-sdk/enums"
)

type errors struct {
	next  abstractLogger
	level enums.LogLevel
}

func (er *errors) setNext(next abstractLogger) {
	er.next = next
}

func (er *errors) setLevel(level enums.LogLevel) {
	er.level = level
}

func (er *errors) Execute(fluentdLogger *FluentdLogger, passedLogLevel enums.LogLevel, tag string, data map[string]string) error {
	var fluentdPostError error
	if er.level == passedLogLevel && passedLogLevel >= enums.GetLogLevelFromLogType(fluentdLogger.InitLogDetails.GlobalLoggingType) {
		fluentdPostError = fluentdLogger.FluentdConnection.Post(tag, data)
	}
	if er.next != nil {
		chainErr := er.next.Execute(fluentdLogger, passedLogLevel, tag, data)
		if chainErr != nil {
			return chainErr
		}
	}

	return fluentdPostError
}
