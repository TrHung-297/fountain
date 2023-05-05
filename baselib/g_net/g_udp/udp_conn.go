/* !!
 * File: udp_conn.go
 * File Created: Thursday, 27th May 2021 10:19:32 am
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Thursday, 27th May 2021 10:35:52 am
 
 */

package gudp

import "net"

// Connection type;
type Connection struct {
	Conn  *net.UDPConn
	Addr  *net.UDPAddr
	codec Codec
}
