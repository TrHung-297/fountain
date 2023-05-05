/* !!
 * File: codec.go
 * File Created: Thursday, 27th May 2021 10:19:32 am
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Thursday, 27th May 2021 10:35:36 am
 
 */

package gudp

import (
	"fmt"
	"io"
)

var (
	protocolRegisters = make(map[string]Protocol)
)

// RegisterProtocol func
func RegisterProtocol(name string, protocol Protocol) {
	protocolRegisters[name] = protocol
}

// NewCodecByName func
func NewCodecByName(name string, rw io.ReadWriter) (Codec, error) {
	protocol, ok := protocolRegisters[name]
	if !ok {
		return nil, fmt.Errorf("KUDP not found protocol name: %s", name)
	}

	return protocol.NewCodec(rw)
}
