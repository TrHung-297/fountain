package base

import "fmt"

func StandardizedNumber(number, lenNumber int) string {
	s := fmt.Sprintf("%d", number)

	if len(s) != lenNumber {
		for i := 0; i < lenNumber-1; i++ {
			s = "0" + s
		}
	}

	return s
}
