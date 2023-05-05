/* !!
 * File: rt.go
 * File Created: Monday, 12th July 2021 6:11:19 pm
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Monday, 12th July 2021 7:24:01 pm
 
 */

package rt

import (
	"io"

	"github.com/TrHung-297/fountain/baselib/g_log/lr"
	"gopkg.in/natefinch/lumberjack.v2"
)

type RotateFileConfig struct {
	Filename   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Level      lr.Level
	Formatter  lr.Formatter
}

type RotateFileHook struct {
	Config    RotateFileConfig
	logWriter io.Writer
}

func NewRotateFileHook(config RotateFileConfig) (lr.Hook, error) {

	hook := RotateFileHook{
		Config: config,
	}
	hook.logWriter = &lumberjack.Logger{
		Filename:   config.Filename,
		MaxSize:    config.MaxSize,
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,
	}

	return &hook, nil
}

func (hook *RotateFileHook) Levels() []lr.Level {
	return lr.AllLevels[:hook.Config.Level+1]
}

func (hook *RotateFileHook) Fire(entry *lr.Entry) (err error) {
	b, err := hook.Config.Formatter.Format(entry)
	if err != nil {
		return err
	}
	hook.logWriter.Write(b)
	return nil
}
