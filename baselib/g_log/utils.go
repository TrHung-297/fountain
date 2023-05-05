package g_log

import "fmt"

func twoDigits(i int) string {
	if i < 10 {
		return fmt.Sprintf("0%d", i)
	}

	return fmt.Sprintf("%d", i)
}
