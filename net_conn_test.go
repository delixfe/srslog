package srslog

import (
	"net"
	"time"
)

type mockConnection struct {
	readFn  func(b []byte) (int, error)
	writeFn func(b []byte) (int, error)
}

func (c *mockConnection) Read(b []byte) (int, error) {
	readFn := c.readFn
	if readFn != nil {
		return readFn(b)
	}
	return len(b), nil
}

func (c *mockConnection) Write(b []byte) (int, error) {
	writeFn := c.writeFn
	if writeFn != nil {
		return writeFn(b)
	}
	return len(b), nil
}
func (c *mockConnection) Close() error { return nil }

func (c *mockConnection) LocalAddr() net.Addr { return &mockAddr{} }

func (c *mockConnection) RemoteAddr() net.Addr { return &mockAddr{} }

func (c *mockConnection) SetDeadline(t time.Time) error { return nil }

func (c *mockConnection) SetReadDeadline(t time.Time) error { return nil }

func (c *mockConnection) SetWriteDeadline(t time.Time) error { return nil }

func (c *mockConnection) SetReadBuffer(bytes int) error { return nil }

func (c *mockConnection) SetWriteBuffer(bytes int) error { return nil }

func mockConnectionFullfillsNetConn() net.Conn {
	return &mockConnection{}
}

type Addr interface {
	Network() string // name of the network (for example, "tcp", "udp")
	String() string  // string form of address (for example, "192.0.2.1:25", "[2001:db8::1]:80")
}

type mockAddr struct {
	network    string
	addrstring string
}

func (m *mockAddr) Network() string {
	return m.network
}
func (m *mockAddr) String() string {
	return m.addrstring
}
