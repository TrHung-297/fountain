/* !!
 * File: example_function_test.go
 * File Created: Monday, 12th July 2021 6:04:09 pm
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Monday, 12th July 2021 7:25:08 pm
 
 */

package lr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogger_LogFn(t *testing.T) {
	SetFormatter(&JSONFormatter{})
	SetLevel(WarnLevel)

	notCalled := 0
	InfoFn(func() []interface{} {
		notCalled++
		return []interface{}{
			"Hello",
		}
	})
	assert.Equal(t, 0, notCalled)

	called := 0
	ErrorFn(func() []interface{} {
		called++
		return []interface{}{
			"Oopsi",
		}
	})
	assert.Equal(t, 1, called)
}
