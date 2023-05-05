package base

import (
	"testing"
)

var dataTest = []string{"2009-01", "2009-01-01", "2009-1", "2009-1-1"}

func TestParseOk(t *testing.T) {
	for _, data := range dataTest {
		timeP, err := TryParseDate(data)
		if err != nil {
			t.Errorf("Parse %q is error %+v", data, err)
		}

		if timeP.Unix() != 1230768000 {
			t.Errorf("Parse %q have result is fail", data)
		}
	}
}
