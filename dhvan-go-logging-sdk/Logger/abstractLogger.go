package Logger

import (
	"dhvan-go-logging-sdk/enums"
)

type abstractLogger interface {
	Execute(*FluentdLogger, enums.LogLevel, string, map[string]string) error
	setNext(abstractLogger)
}
