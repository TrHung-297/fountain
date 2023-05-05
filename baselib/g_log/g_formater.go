package g_log

import (
	"fmt"
	"os"
	"strings"

	"gitlab.gplay.vn/gtv-backend/fountain/baselib/g_log/lr"
)

var pid = os.Getpid()

type GTextFormatter struct {
}

func (GTextFormatter) Format(entry *lr.Entry) ([]byte, error) {
	// Data Fields

	// Time at which the log entry was created
	// Time time.Time

	// Level the log entry was logged at: Trace, Debug, Info, Warn, Error, Fatal or Panic
	// This field will be set on entry firing and the value will be equal to the one in Logger struct field.
	// Level Level

	// Calling method, with package name
	// Caller *runtime.Frame

	// Message passed to Trace, Debug, Info, Warn, Error, Fatal or Panic
	// Message string

	_, month, day := entry.Time.Date()

	// pcs := make([]uintptr, 5)
	// depth := runtime.Callers(5, pcs)
	// frames := runtime.CallersFrames(pcs[:depth])

	frame := entry.Caller
	// for f, again := frames.Next(); again; f, again = frames.Next() {
	// 	pkg := getPackageName(f.Function)

	// 	if strings.EqualFold(pkg, "gitlab.gplay.vn/gtv-backend/fountain/baselib/g_log/lr") || strings.EqualFold(pkg, "gitlab.gplay.vn/gtv-backend/fountain/baselib/g_log") {
	// 		continue
	// 	}

	// 	frame = &f
	// 	break // Important, do not remove
	// }

	pathArr := strings.Split(frame.File, "/")
	path := ""
	if len(pathArr) > 0 {
		path = pathArr[len(pathArr)-1]
	}

	return []byte(fmt.Sprintf("%s%s%s %s   %d %s:%d] %s\n", strings.ToUpper(string(entry.Level.String()[0])), twoDigits(day), twoDigits(int(month)), entry.Time.Format("15:04:05.000000"), pid, path, frame.Line, entry.Message)), nil
}
