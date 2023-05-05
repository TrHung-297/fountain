
package gudp

import (
	"fmt"
	"net"
	"sync/atomic"

	"github.com/TrHung-297/fountain/baselib/g_log"
)

// UDPServer g_udp type;
type UDPServer struct {
	conn      *Connection
	codec     *Codec
	cb        UDPConnectionCallBack
	isRunning *int32
}

// NewUDPServer func;
func NewUDPServer(cb UDPConnectionCallBack) *UDPServer {
	sv := &UDPServer{
		cb:        cb,
		isRunning: new(int32),
	}

	return sv
}

// Serve func;
func (s *UDPServer) Serve() {
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("127.0.0.1:%d", remotePort))
	if err != nil {
		g_log.V(1).WithError(err).Errorf("KUDP Serving ResolveUDPAddr Error: %+v", err)
		return
	}

	conn, err := net.ListenUDP(addr.Network(), addr)
	if err != nil {
		g_log.V(1).WithError(err).Errorf("KUDP ListenUDP Error: %+v", err)
		return
	}

	codec, err := NewCodecByName("g_udp_basic", conn)
	if err != nil {
		g_log.V(1).WithError(err).Errorf("KUDP NewCodecByName Error: %+v", err)
		return
	}

	atomic.AddInt32(s.isRunning, 1)

	udpConn := &Connection{
		Conn:  conn,
		Addr:  addr,
		codec: codec,
	}

	go func() {
		for atomic.LoadInt32(s.isRunning) == 1 {
			msg, err := codec.Receive()
			if err != nil {
				g_log.V(1).WithError(err).Errorf("KUDP Server Receive Error: %+v", err)
				return
			}

			if s.cb != nil {
				s.cb.OnMessageDataArrived(msg)
			}
		}
	}()

	s.conn = udpConn
}

// Close func;
func (s *UDPServer) Close() {
	if s.conn != nil {
		atomic.AddInt32(s.isRunning, 0)
		s.conn.Conn.Close()
		g_log.V(1).Errorf("KUDP Server closed!")
	} else {
		g_log.V(1).Errorf("KUDP Server was closed!")
	}
}
