// +build js nacl plan9

package lr

import (
	"io"
)

func checkIfTerminal(w io.Writer) bool {
	return false
}
