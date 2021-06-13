package tcp

import (
	"fmt"
	"net"
	"time"
)

func Telnet(ip string, port int32) bool {
	var (
		err error
	)

	var conn net.Conn
	if conn, err = net.DialTimeout("tcp", fmt.Sprintf("%s:%d", ip, port), 3*time.Second); err != nil {
		return false
	}

	_ = conn.Close()
	return true
}

func TelnetHost(host string) bool {
	var (
		err error
	)

	var conn net.Conn
	if conn, err = net.DialTimeout("tcp", host, 3*time.Second); err != nil {
		return false
	}

	_ = conn.Close()
	return true
}
