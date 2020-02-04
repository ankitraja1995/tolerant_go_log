package Logger

import (
	"dhvan-go-logging-sdk/enums"
)

func GetChainOfLoggers() abstractLogger {
	warn := &warn{}
	warn.setLevel(enums.WarnLevel)

	error := &errors{}
	error.setLevel(enums.ErrorLevel)

	panic := &panic{}
	panic.setLevel(enums.PanicLevel)

	fatal := &fatal{}
	fatal.setLevel(enums.FatalLevel)

	debug := &debug{}
	debug.setLevel(enums.DebugLevel)

	info := &info{}
	info.setLevel(enums.InfoLevel)

	trace := &trace{}
	trace.setLevel(enums.Trace_level)

	warn.setNext(error)
	error.setNext(panic)
	panic.setNext(fatal)
	fatal.setNext(debug)
	debug.setNext(info)
	info.setNext(trace)

	return warn
}
