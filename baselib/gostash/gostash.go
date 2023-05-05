/* !!
 * File: gostash.go
 * File Created: Wednesday, 5th May 2021 10:37:21 am
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Wednesday, 5th May 2021 10:39:12 am
 
 */

package gostash

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

	"gitlab.gplay.vn/gtv-backend/fountain/baselib/g_log"
)

// Gostash func
type Gostash struct {
	Hostname   string
	Port       int
	Connection *net.TCPConn
	Timeout    int
}

// New func
func New(hostname string, port int, timeout int) *Gostash {
	l := Gostash{}
	l.Hostname = hostname
	l.Port = port
	l.Connection = nil
	l.Timeout = timeout
	return &l
}

func parseURL(url string) (host string, port int, err error) {
	host = "localhost"
	port = 80

	if strings.HasPrefix(url, "https") {
		port = 443
	}

	hostArr := strings.Split(url, "://")
	hostURI := ""
	if len(hostArr) > 1 {
		hostURI = hostArr[1]
	} else {
		hostURI = hostArr[0]
	}
	var lastPos = strings.Index(hostURI, ":")
	if lastPos > -1 {
		host = hostURI[:lastPos]
		p, e := strconv.Atoi(hostURI[lastPos+1:])
		if e != nil {
			err = e
			return
		}
		port = p
	}
	return
}

// NewList func
func NewList(listAddr []string, timeout int) []*Gostash {
	resList := make([]*Gostash, 0)

	for _, addr := range listAddr {
		l := &Gostash{}

		host, port, err := parseURL(addr)
		if err != nil {
			g_log.V(1).WithError(err).Errorf("Address invalid: %s, error: %+v", addr, err)

			continue
		}

		l.Hostname = host
		l.Port = port
		l.Connection = nil
		l.Timeout = timeout

		resList = append(resList, l)
	}

	return resList
}

// Dump func
func (l *Gostash) Dump() {
	fmt.Println("Hostname:   ", l.Hostname)
	fmt.Println("Port:       ", l.Port)
	fmt.Println("Connection: ", l.Connection)
	fmt.Println("Timeout:    ", l.Timeout)
}

// String func
func (l *Gostash) String() string {
	return fmt.Sprintf("Hostname: %s, Port: %d, Timeout: %d", l.Hostname, l.Port, l.Timeout)
}

// SetTimeouts func
func (l *Gostash) SetTimeouts() {
	deadline := time.Now().Add(time.Duration(l.Timeout) * time.Second)
	l.Connection.SetDeadline(deadline)
	l.Connection.SetWriteDeadline(deadline)
	l.Connection.SetReadDeadline(deadline)
}

// Connect func
func (l *Gostash) Connect() (*net.TCPConn, error) {
	var connection *net.TCPConn

	service := fmt.Sprintf("%s:%d", l.Hostname, l.Port)
	addr, err := net.ResolveTCPAddr("tcp", service)
	if err != nil {
		return connection, err
	}

	connection, err = net.DialTCP("tcp", nil, addr)
	if err != nil {
		return connection, err
	}

	if connection != nil {
		l.Connection = connection
		l.Connection.SetLinger(0) // default -1
		l.Connection.SetNoDelay(true)
		l.Connection.SetKeepAlive(true)
		l.Connection.SetKeepAlivePeriod(time.Duration(5) * time.Second)
		l.SetTimeouts()
	}

	return connection, err
}

// Writeln func
func (l *Gostash) Writeln(message string) error {
	var err = errors.New("tcp connection is nil")
	message = fmt.Sprintf("%s\n", message)
	if l.Connection != nil {
		n, err := l.Connection.Write([]byte(message))
		if err != nil {
			g_log.V(1).WithError(err).Errorf("Gostash::Writeln - Error: %+v", err)

			myAddr := l.Connection.RemoteAddr()
			l.Connection.Close()
			l.Connection = nil

			_, conErr := l.Connect()
			g_log.V(1).WithError(err).Errorf("Gostash::Writeln - Re-Connecting to [%+v] conErr: %+v", myAddr, conErr)

			return conErr
		}

		g_log.V(5).Infof("Gostash::Writeln - Connection write n: %d", n)

		// Successful write! Let's extend the timeout.
		l.SetTimeouts()

		return nil
	}

	l.Connect()

	return err
}
