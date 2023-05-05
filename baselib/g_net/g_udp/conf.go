/* !!
 * File: conf.go
 * File Created: Thursday, 27th May 2021 10:19:32 am
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Thursday, 27th May 2021 10:35:40 am
 
 */

package gudp

import (
	"flag"
)

var remotePort int

func init() {
	remotePort = *flag.Int("remote_log_port", 9999, "Config port for remote log udp server")
}
