package logger

const (
	// LogType
	LogTypeWarning = 4
	LogTypeDev     = 5
	// MsgType
	MsgTypeInfo    = 1
	MsgTypeWarning = 2
	MsgTypeError   = 3
)

// SendLogToTelegramBot func
func SendLogToBot(topic, logEvent, message string) {
	// Send log to bot via Inside
}
