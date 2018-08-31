package seelog

import (
	"strings"
	"github.com/xmge/seelog/logs"
)

// Log levels to control the logging output. 提供给外部调用
const (
	LevelAll = iota
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

// BeeLogger references the used application logger.
var SeeLogger = logs.GetSeeLogger()


// listen port
func SetPort(port int)  {
	go server(port)
}

// SetLevel sets the global logs level used by the simple logger.
func SetLevel(l int) {
	logs.SetLevel(l)
}

// SetLogFuncCall set the CallDepth, default is 3
func SetLogFuncCall(b bool) {
	logs.SetLogFuncCall(b)
}

// SetLogger sets a new logger.
func SetLogger(adaptername string, config string) error {
	return logs.SetLogger(adaptername, config)
}

// Critical logs a message at critical level.
func Fatal(v ...interface{}) {
	logs.Fatal(generateFmtStr(len(v)), v...)
}

// Error logs a message at error level.
func Error(v ...interface{}) {
	logs.Error(generateFmtStr(len(v)), v...)
}

// Warn compatibility alias for Warning()
func Warn(v ...interface{}) {
	logs.Warn(generateFmtStr(len(v)), v...)
}

// Info compatibility alias for Warning()
func Info(v ...interface{}) {
	logs.Info(generateFmtStr(len(v)), v...)
}

// Debug logs a message at debug level.
func Debug(v ...interface{}) {
	logs.Debug(generateFmtStr(len(v)), v...)
}

func generateFmtStr(n int) string {
	return strings.Repeat("%v ", n)
}
