/* !!
 * File: inf.go
 * File Created: Thursday, 27th May 2021 10:19:32 am
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Thursday, 27th May 2021 10:35:43 am
 
 */

package gudp

import (
	"io"
)

// Protocol type
type Protocol interface {
	NewCodec(rw io.ReadWriter) (Codec, error)
}

// ProtocolFunc func
type ProtocolFunc func(rw io.ReadWriter) (Codec, error)

// NewCodec func
func (pf ProtocolFunc) NewCodec(rw io.ReadWriter) (Codec, error) {
	return pf(rw)
}

// Codec interface
type Codec interface {
	Receive() (interface{}, error)
	Send(interface{})
	Close()
}

// ConnectionFactory interface
type ConnectionFactory interface {
	NewConnection(serverName string) Connection
}

// ClearSendChan interface
type ClearSendChan interface {
	ClearSendChan(<-chan interface{})
}

// UDPConnectionCallBack func;
type UDPConnectionCallBack interface {
	OnMessageDataArrived(messageRaw interface{})
}
