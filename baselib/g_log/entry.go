package g_log

import (
	"context"
	"time"

	"github.com/TrHung-297/fountain/baselib/g_log/lr"
)

type g_entry struct {
	*lr.Entry
	level uint32
	name  string
}

// WithField allocates a new entry and adds a field to it.
// Debug, Print, Info, Warn, Error, Fatal or Panic must be then applied to
// this new returned e.Entry.
// If you want multiple fields, use `WithFields`.
func (e *g_entry) WithField(key string, value interface{}) *g_entry {
	return &g_entry{Entry: e.Entry.WithField(key, value), level: e.level}
}

// Adds a struct of fields to the log e.Entry. All it does is call `WithField` for
// each `Field`.
func (e *g_entry) WithFields(fields Fields) *g_entry {
	return &g_entry{Entry: e.Entry.WithFields(lr.Fields(fields)), level: e.level}
}

// Add an error as single field to the log e.Entry.  All it does is call
// `WithError` for the given `error`.
func (e *g_entry) WithError(err error) *g_entry {
	return &g_entry{Entry: e.Entry.WithError(err), level: e.level}
}

// Add a context to the log e.Entry.
func (e *g_entry) WithContext(ctx context.Context) *g_entry {
	return &g_entry{Entry: e.Entry.WithContext(ctx), level: e.level}
}

// Overrides the time of the log e.Entry.
func (e *g_entry) WithTime(t time.Time) *g_entry {
	return &g_entry{Entry: e.Entry.WithTime(t), level: e.level}
}

// ----------------------------------------------------------------------------

func (e *g_entry) Log(args ...interface{}) {
	if e.level == 0 || e.level <= logFileLevel || e.level <= logPrintLevel {
		e.Entry.Log(args...)
	}
}

func (e *g_entry) Trace(args ...interface{}) {
	if e.level == 0 || e.level <= logFileLevel || e.level <= logPrintLevel {
		e.Entry.Trace(args...)
	}
}

func (e *g_entry) Debug(args ...interface{}) {
	if e.level == 0 || e.level <= logFileLevel || e.level <= logPrintLevel {
		e.Entry.Debug(args...)
	}
}

func (e *g_entry) Print(args ...interface{}) {
	if e.level == 0 || e.level <= logFileLevel || e.level <= logPrintLevel {
		e.Entry.Print(args...)
	}
}

func (e *g_entry) Info(args ...interface{}) {
	if e.level == 0 || e.level <= logFileLevel || e.level <= logPrintLevel {
		e.Entry.Info(args...)
	}
}

func (e *g_entry) Warn(args ...interface{}) {
	if e.level == 0 || e.level <= logFileLevel || e.level <= logPrintLevel {
		e.Entry.Warn(args...)
	}
}

func (e *g_entry) Warning(args ...interface{}) {
	if e.level == 0 || e.level <= logFileLevel || e.level <= logPrintLevel {
		e.Entry.Warning(args...)
	}
}

func (e *g_entry) Error(args ...interface{}) {
	if e.level == 0 || e.level <= logFileLevel || e.level <= logPrintLevel {
		e.Entry.Error(args...)
	}
}

func (e *g_entry) Fatal(args ...interface{}) {
	if e.level == 0 || e.level <= logFileLevel || e.level <= logPrintLevel {
		e.Entry.Fatal(args...)
	}
}

func (e *g_entry) Panic(args ...interface{}) {
	if e.level == 0 || e.level <= logFileLevel || e.level <= logPrintLevel {
		e.Entry.Panic(args...)
	}
}

// Entry Printf family functions

func (e *g_entry) Logf(level Level, format string, args ...interface{}) {
	if e.level == 0 || level <= Level(logFileLevel) || level <= Level(logPrintLevel) {
		e.Entry.Level = lr.Level(level)
		e.Entry.Logf(format, args...)
	}
}

func (e *g_entry) Tracef(format string, args ...interface{}) {
	if e.level == 0 || e.level <= logFileLevel || e.level <= logPrintLevel {
		e.Entry.Tracef(format, args...)
	}
}

func (e *g_entry) Debugf(format string, args ...interface{}) {
	if e.level == 0 || e.level <= logFileLevel || e.level <= logPrintLevel {
		e.Entry.Debugf(format, args...)
	}
}

func (e *g_entry) Infof(format string, args ...interface{}) {
	if e.level == 0 || e.level <= logFileLevel || e.level <= logPrintLevel {
		e.Entry.Infof(format, args...)
	}
}

func (e *g_entry) Printf(format string, args ...interface{}) {
	if e.level == 0 || e.level <= logFileLevel || e.level <= logPrintLevel {
		e.Entry.Printf(format, args...)
	}
}

func (e *g_entry) Warnf(format string, args ...interface{}) {
	if e.level == 0 || e.level <= logFileLevel || e.level <= logPrintLevel {
		e.Entry.Warnf(format, args...)
	}
}

func (e *g_entry) Warningf(format string, args ...interface{}) {
	if e.level == 0 || e.level <= logFileLevel || e.level <= logPrintLevel {
		e.Entry.Warningf(format, args...)
	}
}

func (e *g_entry) Errorf(format string, args ...interface{}) {
	if e.level == 0 || e.level <= logFileLevel || e.level <= logPrintLevel {
		e.Entry.Errorf(format, args...)
	}
}

func (e *g_entry) Fatalf(format string, args ...interface{}) {
	if e.level == 0 || e.level <= logFileLevel || e.level <= logPrintLevel {
		e.Entry.Fatalf(format, args...)
	}
}

func (e *g_entry) Panicf(format string, args ...interface{}) {
	if e.level == 0 || e.level <= logFileLevel || e.level <= logPrintLevel {
		e.Entry.Panicf(format, args...)
	}
}

// Entry Println family functions

func (e *g_entry) Logln(level Level, args ...interface{}) {
	if e.level == 0 || level <= Level(logFileLevel) || level <= Level(logPrintLevel) {
		e.Entry.Level = lr.Level(level)
		e.Entry.Logln(args...)
	}
}

func (e *g_entry) Traceln(args ...interface{}) {
	if e.level == 0 || e.level <= logFileLevel || e.level <= logPrintLevel {
		e.Entry.Traceln(args...)
	}
}

func (e *g_entry) Debugln(args ...interface{}) {
	if e.level == 0 || e.level <= logFileLevel || e.level <= logPrintLevel {
		e.Entry.Debugln(args...)
	}
}

func (e *g_entry) Infoln(args ...interface{}) {
	if e.level == 0 || e.level <= logFileLevel || e.level <= logPrintLevel {
		e.Entry.Infoln(args...)
	}
}

func (e *g_entry) Println(args ...interface{}) {
	if e.level == 0 || e.level <= logFileLevel || e.level <= logPrintLevel {
		e.Entry.Println(args...)
	}
}

func (e *g_entry) Warnln(args ...interface{}) {
	if e.level == 0 || e.level <= logFileLevel || e.level <= logPrintLevel {
		e.Entry.Warnln(args...)
	}
}

func (e *g_entry) Warningln(args ...interface{}) {
	if e.level == 0 || e.level <= logFileLevel || e.level <= logPrintLevel {
		e.Entry.Warningln(args...)
	}
}

func (e *g_entry) Errorln(args ...interface{}) {
	if e.level == 0 || e.level <= logFileLevel || e.level <= logPrintLevel {
		e.Entry.Errorln(args...)
	}
}

func (e *g_entry) Fatalln(args ...interface{}) {
	if e.level == 0 || e.level <= logFileLevel || e.level <= logPrintLevel {
		e.Entry.Fatalln(args...)
	}
}

func (e *g_entry) Panicln(args ...interface{}) {
	if e.level == 0 || e.level <= logFileLevel || e.level <= logPrintLevel {
		e.Entry.Panicln(args...)
	}
}
