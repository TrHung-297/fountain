// +build appengine

package lr

import (
	"io"
)

func checkIfTerminal(w io.Writer) bool {
	return true
}
