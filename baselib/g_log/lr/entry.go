/* !!
 * File: entry.go
 * File Created: Monday, 12th July 2021 6:04:08 pm
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Monday, 12th July 2021 7:24:55 pm
 
 */

package lr

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"time"
)

var (

	// qualified package name, cached at first use
	lrPackage string

	// Positions in the call stack when tracing to report the calling method
	minimumCallerDepth int

	// Used for caller information initialisation
	callerInitOnce sync.Once
)

const (
	maximumCallerDepth int = 25
	knownLrFrames      int = 5
)

func init() {
	// start at the bottom of the stack before the package-name cache is primed
	minimumCallerDepth = 1
}

// Defines the key when adding errors using WithError.
var ErrorKey = "error"

// An entry is the final or intermediate Lr logging entry. It contains all
// the fields passed with WithField{,s}. It's finally logged when Trace, Debug,
// Info, Warn, Error, Fatal or Panic is called on it. These objects can be
// reused and passed around as much as you wish to avoid field duplication.
type Entry struct {
	Logger *Logger

	// Contains all the fields set by the user.
	Data Fields

	// Time at which the log entry was created
	Time time.Time

	// Level the log entry was logged at: Trace, Debug, Info, Warn, Error, Fatal or Panic
	// This field will be set on entry firing and the value will be equal to the one in Logger struct field.
	Level Level

	// Calling method, with package name
	Caller *runtime.Frame

	// Message passed to Trace, Debug, Info, Warn, Error, Fatal or Panic
	Message string

	// When formatter is called in entry.log(), a Buffer may be set to entry
	Buffer *bytes.Buffer

	// Contains the context set by the user. Useful for hook processing etc.
	Context context.Context

	// err may contain a field formatting error
	err string
}

func NewEntry(logger *Logger) *Entry {
	return &Entry{
		Logger: logger,
		// Default is three fields, plus one optional.  Give a little extra room.
		Data: make(Fields, 6),
	}
}

func (entry *Entry) Dup() *Entry {
	data := make(Fields, len(entry.Data))
	for k, v := range entry.Data {
		data[k] = v
	}
	return &Entry{Logger: entry.Logger, Data: data, Level: entry.Level, Time: entry.Time, Context: entry.Context, err: entry.err}
}

// Returns the bytes representation of this entry from the formatter.
func (entry *Entry) Bytes() ([]byte, error) {
	return entry.Logger.Formatter.Format(entry)
}

// Returns the string representation from the reader and ultimately the
// formatter.
func (entry *Entry) String() (string, error) {
	serialized, err := entry.Bytes()
	if err != nil {
		return "", err
	}
	str := string(serialized)
	return str, nil
}

// Add an error as single field (using the key defined in ErrorKey) to the Entry.
func (entry *Entry) WithError(err error) *Entry {
	return entry.WithField(ErrorKey, err)
}

// Add a context to the Entry.
func (entry *Entry) WithContext(ctx context.Context) *Entry {
	dataCopy := make(Fields, len(entry.Data))
	for k, v := range entry.Data {
		dataCopy[k] = v
	}
	return &Entry{Logger: entry.Logger, Data: dataCopy, Level: entry.Level, Time: entry.Time, err: entry.err, Context: ctx}
}

// Add a single field to the Entry.
func (entry *Entry) WithField(key string, value interface{}) *Entry {
	return entry.WithFields(Fields{key: value})
}

// Add a map of fields to the Entry.
func (entry *Entry) WithFields(fields Fields) *Entry {
	data := make(Fields, len(entry.Data)+len(fields))
	for k, v := range entry.Data {
		data[k] = v
	}
	fieldErr := entry.err
	for k, v := range fields {
		isErrField := false
		if t := reflect.TypeOf(v); t != nil {
			switch {
			case t.Kind() == reflect.Func, t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Func:
				isErrField = true
			}
		}
		if isErrField {
			tmp := fmt.Sprintf("can not add field %q", k)
			if fieldErr != "" {
				fieldErr = entry.err + ", " + tmp
			} else {
				fieldErr = tmp
			}
		} else {
			data[k] = v
		}
	}
	return &Entry{Logger: entry.Logger, Data: data, Level: entry.Level, Time: entry.Time, err: fieldErr, Context: entry.Context}
}

// Overrides the time of the Entry.
func (entry *Entry) WithTime(t time.Time) *Entry {
	dataCopy := make(Fields, len(entry.Data))
	for k, v := range entry.Data {
		dataCopy[k] = v
	}
	return &Entry{Logger: entry.Logger, Data: dataCopy, Level: entry.Level, Time: t, err: entry.err, Context: entry.Context}
}

// getPackageName reduces a fully qualified function name to the package name
// There really ought to be to be a better way...
func getPackageName(f string) string {
	for {
		lastPeriod := strings.LastIndex(f, ".")
		lastSlash := strings.LastIndex(f, "/")
		if lastPeriod > lastSlash {
			f = f[:lastPeriod]
		} else {
			break
		}
	}

	return f
}

// getCaller retrieves the name of the first non-lr calling function
func getCaller() *runtime.Frame {
	// cache this package's fully-qualified name
	callerInitOnce.Do(func() {
		pcs := make([]uintptr, maximumCallerDepth)
		_ = runtime.Callers(0, pcs)

		// dynamic get the package name and the minimum caller depth
		for i := 0; i < maximumCallerDepth; i++ {
			funcName := runtime.FuncForPC(pcs[i]).Name()
			if strings.Contains(funcName, "getCaller") {
				lrPackageRaw := getPackageName(funcName)
				lastSplatIdx := strings.LastIndex(lrPackageRaw, "/")
				glogPackage := lrPackageRaw[:lastSplatIdx]
				lrPackage = glogPackage
				break
			}
		}

		minimumCallerDepth = knownLrFrames
	})

	// Restrict the lookback frames to avoid runaway lookups
	pcs := make([]uintptr, maximumCallerDepth)
	depth := runtime.Callers(minimumCallerDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])

	for f, again := frames.Next(); again; f, again = frames.Next() {
		pkg := getPackageName(f.Function)

		// If the caller isn't part of this package, we're done
		if !strings.Contains(pkg, lrPackage) {
			return &f //nolint:scopelint
		}
	}

	// if we got here, we failed to find the caller's context
	return nil
}

func (entry Entry) HasCaller() (has bool) {
	return entry.Logger != nil &&
		entry.Logger.ReportCaller &&
		entry.Caller != nil
}

func (entry *Entry) log(msg string) {
	var buffer *bytes.Buffer

	newEntry := entry.Dup()

	if newEntry.Time.IsZero() {
		newEntry.Time = time.Now()
	}

	newEntry.Level = entry.Level
	newEntry.Message = msg

	newEntry.Logger.mu.Lock()
	reportCaller := newEntry.Logger.ReportCaller
	newEntry.Logger.mu.Unlock()

	if reportCaller {
		newEntry.Caller = getCaller()
	}

	newEntry.fireHooks()

	buffer = getBuffer()
	defer func() {
		newEntry.Buffer = nil
		putBuffer(buffer)
	}()
	buffer.Reset()
	newEntry.Buffer = buffer

	// If not print to stdout
	newEntry.write()

	newEntry.Buffer = nil
}

func (entry *Entry) fireHooks() {
	var tmpHooks LevelHooks
	entry.Logger.mu.Lock()
	tmpHooks = make(LevelHooks, len(entry.Logger.Hooks))
	for k, v := range entry.Logger.Hooks {
		tmpHooks[k] = v
	}
	entry.Logger.mu.Unlock()

	err := tmpHooks.Fire(entry.Level, entry)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to fire hook: %v\n", err)
	}
}

func (entry *Entry) write() {
	if !entry.Logger.IsLevelEnabled(entry.Level) {
		return
	}

	serialized, err := entry.Logger.Formatter.Format(entry)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to obtain reader, %v\n", err)
		return
	}
	entry.Logger.mu.Lock()
	defer entry.Logger.mu.Unlock()
	if _, err := entry.Logger.Out.Write(serialized); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to write to log, %v\n", err)
	}
}

func (entry *Entry) Log(args ...interface{}) {
	entry.log(fmt.Sprint(args...))
}

func (entry *Entry) Trace(args ...interface{}) {
	entry.WithField("log_type", "trace").Log(args...)
}

func (entry *Entry) Debug(args ...interface{}) {
	entry.WithField("log_type", "debug").Log(args...)
}

func (entry *Entry) Print(args ...interface{}) {
	entry.Info(args...)
}

func (entry *Entry) Info(args ...interface{}) {
	entry.WithField("log_type", "info").Log(args...)
}

func (entry *Entry) Warn(args ...interface{}) {
	entry.WithField("log_type", "warning").Log(args...)
}

func (entry *Entry) Warning(args ...interface{}) {
	entry.Warn(args...)
}

func (entry *Entry) Error(args ...interface{}) {
	entry.WithField("log_type", "error").Log(args...)
}

func (entry *Entry) Fatal(args ...interface{}) {
	entry.WithField("log_type", "fatal").Log(args...)
	entry.Logger.Exit(1)
}

func (entry *Entry) Panic(args ...interface{}) {
	entry.WithField("log_type", "panic").Log(args...)
}

// Entry Printf family functions

func (entry *Entry) Logf(format string, args ...interface{}) {
	entry.Log(fmt.Sprintf(format, args...))
}

func (entry *Entry) Tracef(format string, args ...interface{}) {
	entry.WithField("log_type", "trace").Logf(format, args...)
}

func (entry *Entry) Debugf(format string, args ...interface{}) {
	entry.WithField("log_type", "debug").Logf(format, args...)
}

func (entry *Entry) Infof(format string, args ...interface{}) {
	entry.WithField("log_type", "info").Logf(format, args...)
}

func (entry *Entry) Printf(format string, args ...interface{}) {
	entry.Infof(format, args...)
}

func (entry *Entry) Warnf(format string, args ...interface{}) {
	entry.WithField("log_type", "warning").Logf(format, args...)
}

func (entry *Entry) Warningf(format string, args ...interface{}) {
	entry.Warnf(format, args...)
}

func (entry *Entry) Errorf(format string, args ...interface{}) {
	entry.WithField("log_type", "error").Logf(format, args...)
}

func (entry *Entry) Fatalf(format string, args ...interface{}) {
	entry.WithField("log_type", "fatal").Logf(format, args...)
	entry.Logger.Exit(1)
}

func (entry *Entry) Panicf(format string, args ...interface{}) {
	entry.WithField("log_type", "panic").Logf(format, args...)
}

// Entry Println family functions

func (entry *Entry) Logln(args ...interface{}) {
	entry.Log(entry.sprintlnn(args...))
}

func (entry *Entry) Traceln(args ...interface{}) {
	entry.WithField("log_type", "trace").Logln(args...)
}

func (entry *Entry) Debugln(args ...interface{}) {
	entry.WithField("log_type", "debug").Logln(args...)
}

func (entry *Entry) Infoln(args ...interface{}) {
	entry.WithField("log_type", "info").Logln(args...)
}

func (entry *Entry) Println(args ...interface{}) {
	entry.Infoln(args...)
}

func (entry *Entry) Warnln(args ...interface{}) {
	entry.WithField("log_type", "warning").Logln(args...)
}

func (entry *Entry) Warningln(args ...interface{}) {
	entry.Warnln(args...)
}

func (entry *Entry) Errorln(args ...interface{}) {
	entry.WithField("log_type", "error").Logln(args...)
}

func (entry *Entry) Fatalln(args ...interface{}) {
	entry.WithField("log_type", "fatal").Logln(args...)
	entry.Logger.Exit(1)
}

func (entry *Entry) Panicln(args ...interface{}) {
	entry.WithField("log_type", "panic").Logln(args...)
}

// Sprintlnn => Sprint no newline. This is to get the behavior of how
// fmt.Sprintln where spaces are always added between operands, regardless of
// their type. Instead of vendoring the Sprintln implementation to spare a
// string allocation, we do the simplest thing.
func (entry *Entry) sprintlnn(args ...interface{}) string {
	msg := fmt.Sprintln(args...)
	return msg[:len(msg)-1]
}
