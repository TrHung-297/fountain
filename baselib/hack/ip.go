

package hack

import (
	"errors"
	"net"
)

var externalIP string

// ExternalIP func;
func ExternalIP() (string, error) {
	if externalIP != "" {
		return externalIP, nil
	}

	infList, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, inf := range infList {
		if inf.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if inf.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := inf.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}

			externalIP = ip.String()
			return externalIP, nil
		}
	}
	return "", errors.New("are you connected to the network?")
}
