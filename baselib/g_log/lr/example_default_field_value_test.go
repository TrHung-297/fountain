/* !!
 * File: example_default_field_value_test.go
 * File Created: Monday, 12th July 2021 6:04:09 pm
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Monday, 12th July 2021 7:25:02 pm
 
 */

package lr

import (
	"os"
)

type DefaultFieldHook struct {
	GetValue func() string
}

func (h *DefaultFieldHook) Levels() []Level {
	return AllLevels
}

func (h *DefaultFieldHook) Fire(e *Entry) error {
	e.Data["aDefaultField"] = h.GetValue()
	return nil
}

func ExampleDefaultFieldHook() {
	l := New()
	l.Out = os.Stdout
	l.Formatter = &TextFormatter{DisableTimestamp: true, DisableColors: true}

	l.AddHook(&DefaultFieldHook{GetValue: func() string { return "with its default value" }})
	l.Info("first log")
	// Output:
	// level=info msg="first log" aDefaultField="with its default value"
}
