/* !!
 * File: testutils.go
 * File Created: Monday, 12th July 2021 6:04:09 pm
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Monday, 12th July 2021 7:26:56 pm
 
 */

package lr

import (
	"bytes"
	"encoding/json"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func logAndAssertJSON(t *testing.T, log func(*Logger), assertions func(fields Fields)) {
	var buffer bytes.Buffer
	var fields Fields

	logger := New()
	logger.Out = &buffer
	logger.Formatter = new(JSONFormatter)

	log(logger)

	err := json.Unmarshal(buffer.Bytes(), &fields)
	require.Nil(t, err)

	assertions(fields)
}

func logAndAssertText(t *testing.T, log func(*Logger), assertions func(fields map[string]string)) {
	var buffer bytes.Buffer

	logger := New()
	logger.Out = &buffer
	logger.Formatter = &TextFormatter{
		DisableColors: true,
	}

	log(logger)

	fields := make(map[string]string)
	for _, kv := range strings.Split(strings.TrimRight(buffer.String(), "\n"), " ") {
		if !strings.Contains(kv, "=") {
			continue
		}
		kvArr := strings.Split(kv, "=")
		key := strings.TrimSpace(kvArr[0])
		val := kvArr[1]
		if kvArr[1][0] == '"' {
			var err error
			val, err = strconv.Unquote(val)
			require.NoError(t, err)
		}
		fields[key] = val
	}
	assertions(fields)
}
