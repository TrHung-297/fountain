package lr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFieldValueError(t *testing.T) {
	buf := &bytes.Buffer{}
	l := &Logger{
		Out:       buf,
		Formatter: new(JSONFormatter),
		Hooks:     make(LevelHooks),
		Level:     DebugLevel,
	}
	l.WithField("func", func() {}).Info("test")
	fmt.Println(buf.String())
	var data map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &data); err != nil {
		t.Error("unexpected error", err)
	}
	_, ok := data[FieldKeyLrError]
	require.True(t, ok, `cannot found expected "lr_error" field: %v`, data)
}

func TestNoFieldValueError(t *testing.T) {
	buf := &bytes.Buffer{}
	l := &Logger{
		Out:       buf,
		Formatter: new(JSONFormatter),
		Hooks:     make(LevelHooks),
		Level:     DebugLevel,
	}
	l.WithField("str", "str").Info("test")
	fmt.Println(buf.String())
	var data map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &data); err != nil {
		t.Error("unexpected error", err)
	}
	_, ok := data[FieldKeyLrError]
	require.False(t, ok)
}

func TestWarninglnNotEqualToWarning(t *testing.T) {
	buf := &bytes.Buffer{}
	bufln := &bytes.Buffer{}

	formatter := new(TextFormatter)
	formatter.DisableTimestamp = true
	formatter.DisableLevelTruncation = true

	l := &Logger{
		Out:       buf,
		Formatter: formatter,
		Hooks:     make(LevelHooks),
		Level:     DebugLevel,
	}
	l.Warning("hello,", "world")

	l.SetOutput(bufln)
	l.Warningln("hello,", "world")

	assert.NotEqual(t, buf.String(), bufln.String(), "Warning() and Wantingln() should not be equal")
}
