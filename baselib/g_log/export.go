package g_log

import (
	"context"
	"time"

	"gitlab.gplay.vn/gtv-backend/fountain/baselib/g_log/lr"
)

// L func set level of log for g_log instance
func L(level uint32) *g_entry {
	log := newGLog()
	log.level = level

	entry := log.logger.WithField("v", level)
	entry.Level = lr.Level(level)
	return &g_entry{Entry: entry, level: level}
}

// V func set level of log for g_log instance
func V(level uint32) *g_entry {
	log := newGLog()
	log.level = level

	entry := log.logger.WithField("v", level)
	entry.Level = lr.Level(level)
	return &g_entry{Entry: entry, level: level}
}

// N func set name of log for g_log instance
func (kg *g_log) N(name string) *g_entry {
	kg.name = name
	return &g_entry{Entry: kg.logger.WithField("name", name), name: name}
}

// C func set context of log for g_log instance
func (kg *g_log) C(c context.Context) *g_entry {
	kg.ctx = c
	return &g_entry{Entry: kg.logger.WithContext(c)}
}

// ----------------------------------------------------------------------

// WithField allocates a new entry and adds a field to it.
// Debug, Print, Info, Warn, Error, Fatal or Panic must be then applied to
// this new returned entry.
// If you want multiple fields, use `WithFields`.
func WithField(key string, value interface{}) *g_entry {
	return newGLog().WithField(key, value)
}

// Adds a struct of fields to the log entry. All it does is call `WithField` for
// each `Field`.
func WithFields(fields Fields) *g_entry {
	return newGLog().WithFields(fields)
}

// Add an error as single field to the log entry.  All it does is call
// `WithError` for the given `error`.
func WithError(err error) *g_entry {
	return newGLog().WithError(err)
}

// Add a context to the log entry.
func WithContext(ctx context.Context) *g_entry {
	return newGLog().WithContext(ctx)
}

// Overrides the time of the log entry.
func WithTime(t time.Time) *g_entry {
	return newGLog().WithTime(t)
}

func Logf(level Level, format string, args ...interface{}) {
	newGLog().logger.Logf(level.ToLrLevel(), format, args...)
}

func Tracef(format string, args ...interface{}) {
	newGLog().logger.Tracef(format, args...)
}

func Debugf(format string, args ...interface{}) {
	newGLog().logger.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	newGLog().logger.Infof(format, args...)
}

func Printf(format string, args ...interface{}) {
	newGLog().logger.Printf(format, args...)
}

func Warnf(format string, args ...interface{}) {
	newGLog().logger.Warnf(format, args...)
}

func Warningf(format string, args ...interface{}) {
	newGLog().logger.Warningf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	newGLog().logger.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	newGLog().logger.Fatalf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	newGLog().logger.Panicf(format, args...)
}

func Log(level Level, args ...interface{}) {
	newGLog().logger.Log(level.ToLrLevel(), args...)
}

func LogFn(level Level, fn LogFunction) {
	newGLog().logger.LogFn(level.ToLrLevel(), fn.ToLrFunction())
}

func Trace(args ...interface{}) {
	newGLog().logger.Trace(args...)
}

func Debug(args ...interface{}) {
	newGLog().logger.Debug(args...)
}

func Info(args ...interface{}) {
	newGLog().logger.Info(args...)
}

func Print(args ...interface{}) {
	newGLog().logger.Print(args...)
}

func Warn(args ...interface{}) {
	newGLog().logger.Warn(args...)
}

func Warning(args ...interface{}) {
	newGLog().logger.Warning(args...)
}

func Error(args ...interface{}) {
	newGLog().logger.Error(args...)
}

func Fatal(args ...interface{}) {
	newGLog().logger.Fatal(args...)
}

func Panic(args ...interface{}) {
	newGLog().logger.Panic(args...)
}

func TraceFn(fn LogFunction) {
	newGLog().logger.TraceFn(fn.ToLrFunction())
}

func DebugFn(fn LogFunction) {
	newGLog().logger.DebugFn(fn.ToLrFunction())
}

func InfoFn(fn LogFunction) {
	newGLog().logger.InfoFn(fn.ToLrFunction())
}

func PrintFn(fn LogFunction) {
	newGLog().logger.PrintFn(fn.ToLrFunction())
}

func WarnFn(fn LogFunction) {
	newGLog().logger.WarnFn(fn.ToLrFunction())
}

func WarningFn(fn LogFunction) {
	newGLog().logger.WarningFn(fn.ToLrFunction())
}

func ErrorFn(fn LogFunction) {
	newGLog().logger.ErrorFn(fn.ToLrFunction())
}

func FatalFn(fn LogFunction) {
	newGLog().logger.FatalFn(fn.ToLrFunction())
}

func PanicFn(fn LogFunction) {
	newGLog().logger.PanicFn(fn.ToLrFunction())
}

func Logln(level Level, args ...interface{}) {
	newGLog().logger.Logln(level.ToLrLevel(), args...)
}

func Traceln(args ...interface{}) {
	newGLog().logger.Traceln(args...)
}

func Debugln(args ...interface{}) {
	newGLog().logger.Debugln(args...)
}

func Infoln(args ...interface{}) {
	newGLog().logger.Infoln(args...)
}

func Println(args ...interface{}) {
	newGLog().logger.Println(args...)
}

func Warnln(args ...interface{}) {
	newGLog().logger.Warnln(args...)
}

func Warningln(args ...interface{}) {
	newGLog().logger.Warningln(args...)
}

func Errorln(args ...interface{}) {
	newGLog().logger.Errorln(args...)
}

func Fatalln(args ...interface{}) {
	newGLog().logger.Fatalln(args...)
}

func Panicln(args ...interface{}) {
	newGLog().logger.Panicln(args...)
}
