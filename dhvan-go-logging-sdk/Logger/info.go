package Logger

import (
	"dhvan-go-logging-sdk/enums"
)

type info struct {
	next  abstractLogger
	level enums.LogLevel
}

func (inf *info) setNext(next abstractLogger) {
	inf.next = next
}

func (inf *info) setLevel(level enums.LogLevel) {
	inf.level = level
}

func (inf *info) Execute(fluentdLogger *FluentdLogger, passedLogLevel enums.LogLevel, tag string, data map[string]string) error {
	var fluentdPostError error
	if inf.level == passedLogLevel && passedLogLevel >= enums.GetLogLevelFromLogType(fluentdLogger.InitLogDetails.GlobalLoggingType) {
		fluentdPostError = fluentdLogger.FluentdConnection.Post(tag, data)
	}
	if inf.next != nil {
		chainErr := inf.next.Execute(fluentdLogger, passedLogLevel, tag, data)
		if chainErr != nil {
			return chainErr
		}
	}

	return fluentdPostError
}
