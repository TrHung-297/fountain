package g_log

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mattn/go-colorable"
	"gitlab.gplay.vn/gtv-backend/fountain/baselib/g_log/lr"
	"gitlab.gplay.vn/gtv-backend/fountain/baselib/g_log/rt"
)

var logGInstance *lr.Logger

type g_log struct {
	logger *lr.Logger
	level  uint32
	name   string
	ctx    context.Context
}

func newGLog() *g_log {
	if logGInstance != nil {
		return &g_log{logger: logGInstance}
	}

	logInstance := lr.New()
	// The API for setting attributes is a little different than the package level
	// exported logger. See Godoc.
	logInstance.Out = os.Stdout

	logInstance.SetReportCaller(true)
	// log.SetFormatter(&lr.JSONFormatter{})

	// You could set this to any `io.Writer` such as a file
	// file, err := os.OpenFile("lr.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// if err == nil {
	// 	log.Out = file
	// } else {
	// 	log.Info("Failed to log to file, using default stderr")
	// }

	serviceIndentity := fmt.Sprintf("%s_%s_%s_%s_log.log", serviceName, hostName, podName, serverID)
	for strings.Contains(serviceIndentity, "__") {
		serviceIndentity = strings.ReplaceAll(serviceIndentity, "__", "_")
	}

	rotateFileHook, err := rt.NewRotateFileHook(rt.RotateFileConfig{
		Filename:   filepath.Join(logDir, strings.Trim(serviceIndentity, "_")),
		MaxSize:    1, // megabytes
		MaxBackups: 3,
		MaxAge:     28, //days
		Level:      lr.Level(logFileLevel),
		Formatter: &lr.JSONFormatter{
			TimestampFormat: "02-01-2006 15:04:05.000000",
			// PrettyPrint:     true,
			DisableHTMLEscape: true,
		},
	})

	if err != nil {
		logInstance.Fatalf("Failed to initialize file rotate hook: %v", err)
	}

	logInstance.SetLevel(lr.Level(logPrintLevel))
	logInstance.SetOutput(colorable.NewColorableStdout())

	// log.SetFormatter(&lr.TextFormatter{
	// 	ForceColors:     true,
	// 	FullTimestamp:   true,
	// 	TimestampFormat: time.RFC822,
	// })
	logInstance.SetFormatter(new(GTextFormatter))

	logInstance.AddHook(rotateFileHook)
	logGInstance = logInstance

	return &g_log{logger: logInstance}
}
