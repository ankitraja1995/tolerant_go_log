package Logger

import (
	"dhvan-go-logging-sdk/enums"
)

type warn struct {
	next  abstractLogger
	level enums.LogLevel
}

func (wrn *warn) setNext(next abstractLogger) {
	wrn.next = next
}

func (wrn *warn) setLevel(level enums.LogLevel) {
	wrn.level = level
}

func (wrn *warn) Execute(fluentdLogger *FluentdLogger, passedLogLevel enums.LogLevel, tag string, data map[string]string) error {
	var fluentdPostError error
	if wrn.level == passedLogLevel && passedLogLevel >= enums.GetLogLevelFromLogType(fluentdLogger.InitLogDetails.GlobalLoggingType) {
		fluentdPostError = fluentdLogger.FluentdConnection.Post(tag, data)
	}
	if wrn.next != nil {
		chainErr := wrn.next.Execute(fluentdLogger, passedLogLevel, tag, data)
		if chainErr != nil {
			return chainErr
		}
	}

	return fluentdPostError
}
