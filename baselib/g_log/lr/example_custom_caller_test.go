package lr

import (
	"os"
	"path"
	"runtime"
	"strings"
)

func ExampleJSONFormatter_CallerPrettyfier() {
	l := New()
	l.SetReportCaller(true)
	l.Out = os.Stdout
	l.Formatter = &JSONFormatter{
		DisableTimestamp: true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			s := strings.Split(f.Function, ".")
			funcname := s[len(s)-1]
			_, filename := path.Split(f.File)
			return funcname, filename
		},
	}
	l.Info("example of custom format caller")
	// Output:
	// {"file":"example_custom_caller_test.go","func":"ExampleJSONFormatter_CallerPrettyfier","level":"info","msg":"example of custom format caller"}
}
