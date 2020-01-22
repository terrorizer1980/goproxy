package goproxy

import (
	"net"
	"time"
)

type TimeoutConn struct {
	net.Conn
	timeout time.Duration
}

func NewTimeoutConn(timeout time.Duration, conn net.Conn) net.Conn {
	tc := &TimeoutConn{
		Conn:    conn,
		timeout: timeout,
	}

	conn.SetDeadline(time.Now().Add(timeout))
	return tc
}

func (tc *TimeoutConn) Read(b []byte) (n int, err error) {
	tc.Conn.SetDeadline(time.Now().Add(tc.timeout))
	return tc.Conn.Read(b)
}

func (tc *TimeoutConn) Write(b []byte) (n int, err error) {
	tc.Conn.SetDeadline(time.Now().Add(tc.timeout))
	return tc.Conn.Write(b)
}
