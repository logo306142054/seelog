package seelog

import (
	"strings"
)

// Log levels to control the logging output.
const (
	LevelAll = "all"
	LevelTrace = "trace"
	LevelDebug = "debug"
	LevelInfo = "info"
	LevelWarn = "warn"
	LevelError = "error"
	LevelFatal = "fatal"
	LevelOff = "off"
)

// log output target
const (
	TargetConsole = iota
	TargetBrowser
	TargetFile
	TargetMail
)

type log interface {
	Debug()
}

// BeeLogger references the used application logger.
var BeeLogger = logs.GetBeeLogger()

// SetLevel sets the global log level used by the simple logger.
func SetLevel(l int) {
	logs.SetLevel(l)
}

// SetLogger sets a new logger.
func SetLogger(adaptername string, config string) error {
	return logs.SetLogger(adaptername, config)
}

// Fatal logs a message at error level.
func Trace(v ...interface{}) {
	logs.Error(generateFmtStr(len(v)), v...)
}

// Error logs a message at error level.
func Debug(v ...interface{}) {
	logs.Error(generateFmtStr(len(v)), v...)
}

// Warning logs a message at warning level.
func Info(v ...interface{}) {
	logs.Warning(generateFmtStr(len(v)), v...)
}

// Warn compatibility alias for Warning()
func Error(v ...interface{}) {
	logs.Warn(generateFmtStr(len(v)), v...)
}

// Notice logs a message at notice level.
func Fatal(v ...interface{}) {
	logs.Notice(generateFmtStr(len(v)), v...)
}

func generateFmtStr(n int) string {
	return strings.Repeat("%v ", n)
}
