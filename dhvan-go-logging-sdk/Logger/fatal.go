package Logger

import (
	"dhvan-go-logging-sdk/enums"
)

type fatal struct {
	next  abstractLogger
	level enums.LogLevel
}

func (fatl *fatal) setNext(next abstractLogger) {
	fatl.next = next
}

func (fatl *fatal) setLevel(level enums.LogLevel) {
	fatl.level = level
}

func (fatl *fatal) Execute(fluentdLogger *FluentdLogger, passedLogLevel enums.LogLevel, tag string, data map[string]string) error {
	var fluentdPostError error
	if fatl.level == passedLogLevel && passedLogLevel >= enums.GetLogLevelFromLogType(fluentdLogger.InitLogDetails.GlobalLoggingType) {
		fluentdPostError = fluentdLogger.FluentdConnection.Post(tag, data)
	}
	if fatl.next != nil {
		chainErr := fatl.next.Execute(fluentdLogger, passedLogLevel, tag, data)
		if chainErr != nil {
			return chainErr
		}
	}

	return fluentdPostError
}
