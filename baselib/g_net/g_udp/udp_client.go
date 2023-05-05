
package gudp

import (
	"fmt"
	"net"

	"github.com/TrHung-297/fountain/baselib/g_log"
)

// UDPClient type;
type UDPClient struct {
	conn  *Connection
	codec *Codec
}

var clientInstance *UDPClient

// NewKUDPClient func;
func NewKUDPClient() *UDPClient {
	if clientInstance != nil {
		return clientInstance
	}

	clientInstance = &UDPClient{}

	return clientInstance
}

// Connect func
func (c *UDPClient) Connect() {
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("127.0.0.1:%d", remotePort))
	if err != nil {
		g_log.V(1).WithError(err).Errorf("KUDP Connect ResolveUDPAddr Error: %+v", err)
		return
	}

	conn, err := net.DialUDP(addr.Network(), nil, addr)
	if err != nil {
		g_log.V(1).WithError(err).Errorf("KUDP DialUDP Error: %+v", err)
		return
	}

	codec, err := NewCodecByName("g_udp_basic", conn)
	if err != nil {
		g_log.V(1).WithError(err).Errorf("KUDP NewCodecByName Error: %+v", err)
		return
	}

	udpConn := &Connection{
		Conn:  conn,
		Addr:  addr,
		codec: codec,
	}

	c.conn = udpConn
}

// SendMessage func
func (c *UDPClient) SendMessage(message interface{}) {
	c.conn.codec.Send(message)
}

// Close func;
func (c *UDPClient) Close() {
	if c.conn != nil {
		c.conn.Conn.Close()
		g_log.V(1).Error("KUDP Client closed!")
	} else {
		g_log.V(1).Error("KUDP Client was closed!")
	}
}
