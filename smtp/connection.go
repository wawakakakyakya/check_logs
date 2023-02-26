package smtp

import (
	"crypto/tls"
	"fmt"
	"net"
)

func plainConnection(addr string, port int) (*net.Conn, error) {
	server := fmt.Sprintf("%s:%d", addr, port)

	conn, err := net.Dial("tcp", server)
	if err != nil {
		return nil, err
	}
	return &conn, nil
}

func tlsConnection(addr string, port int) (*tls.Conn, error) {
	server := fmt.Sprintf("%s:%d", addr, port)
	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         addr,
	}

	// TLSで通信するためのコネクションを用意する
	con, err := tls.Dial("tcp", server, tlsconfig)
	if err != nil {
		return nil, err
	}
	return con, nil
}
