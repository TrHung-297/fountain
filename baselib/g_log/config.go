package g_log

import (
	"gitlab.gplay.vn/gtv-backend/fountain/baselib/g_log/lr"
)

type Fields map[string]interface{}

func (f Fields) ToLrFields() lr.Fields {
	return lr.Fields(f)
}

type Level uint32

func (t Level) ToLrLevel() lr.Level {
	return lr.Level(t - 1)
}

// These are the different logging levels. You can set the logging level to log
// on your instance of logger, obtained with `g_log.New()`.
const (
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	PanicLevel Level = iota + 1
	// FatalLevel level. Logs and then calls `g_log.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
	// TraceLevel level. Designates finer-grained informational events than the Debug.
	TraceLevel
)

// LogFunction For big messages, it can be more efficient to pass a function
// and only call it if the log level is actually enables rather than
// generating the log message and then checking if the level is enabled
type LogFunction func() []interface{}

func (f LogFunction) ToLrFunction() lr.LogFunction {
	return lr.LogFunction(f)
}

var (
	enableRemoteLog bool
	enableTracing   bool
	logFileLevel    uint32
	logPrintLevel   uint32
	hostName        string
	podName         string
	serviceName     string
	serverName      string
	serverID        string
	logDir          string
)

// EnableRemoteLog func;
func EnableRemoteLog(enable bool) {
	enableRemoteLog = enable
}

// EnableTracing func;
func EnableTracing(enable bool) {
	enableTracing = enable
}

// LogDir func;
func LogDir(dir string) {
	logDir = dir
}

// LogFileLevel set log level will be print to file;
func LogFileLevel(level uint32) {
	logFileLevel = level

}

// LogPrintLevel set log level will be print to stdout;
func LogPrintLevel(level uint32) {
	logPrintLevel = level
}

// HostName func;
func HostName(name string) {
	hostName = name
}

// PodName func;
func PodName(name string) {
	podName = name
}

// ServiceName func;
func ServiceName(svName string) {
	serviceName = svName
}

// ServerName func;
func ServerName(sv string) {
	serverName = sv
}

// ServerID func;
func ServerID(id string) {
	serverID = id
}
