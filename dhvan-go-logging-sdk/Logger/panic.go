package Logger

import (
	"dhvan-go-logging-sdk/enums"
)

type panic struct {
	next  abstractLogger
	level enums.LogLevel
}

func (pnic *panic) setNext(next abstractLogger) {
	pnic.next = next
}

func (pnic *panic) setLevel(level enums.LogLevel) {
	pnic.level = level
}

func (pnic *panic) Execute(fluentdLogger *FluentdLogger, passedLogLevel enums.LogLevel, tag string, data map[string]string) error {
	var fluentdPostError error
	if pnic.level == passedLogLevel && passedLogLevel >= enums.GetLogLevelFromLogType(fluentdLogger.InitLogDetails.GlobalLoggingType) {
		fluentdPostError = fluentdLogger.FluentdConnection.Post(tag, data)
	}
	if pnic.next != nil {
		chainErr := pnic.next.Execute(fluentdLogger, passedLogLevel, tag, data)
		if chainErr != nil {
			return chainErr
		}
	}

	return fluentdPostError
}
