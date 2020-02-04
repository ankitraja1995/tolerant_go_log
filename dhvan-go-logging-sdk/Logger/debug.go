package Logger

import (
	"dhvan-go-logging-sdk/enums"
)

type debug struct {
	next  abstractLogger
	level enums.LogLevel
}

func (dbg *debug) setNext(next abstractLogger) {
	dbg.next = next
}

func (dbg *debug) setLevel(level enums.LogLevel) {
	dbg.level = level
}

func (dbg *debug) Execute(fluentdLogger *FluentdLogger, passedLogLevel enums.LogLevel, tag string, data map[string]string) error {
	var fluentdPostError error
	if dbg.level == passedLogLevel && passedLogLevel >= enums.GetLogLevelFromLogType(fluentdLogger.InitLogDetails.GlobalLoggingType) {
		fluentdPostError = fluentdLogger.FluentdConnection.Post(tag, data)
	}
	if dbg.next != nil {
		chainErr := dbg.next.Execute(fluentdLogger, passedLogLevel, tag, data)
		if chainErr != nil {
			return chainErr
		}
	}

	return fluentdPostError
}
