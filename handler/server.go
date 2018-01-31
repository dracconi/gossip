package handler

import (
	"crypto/cipher"
	"net"
)

// Server struct
type Server struct {
	Conn  net.Conn
	Conns []net.Conn

	cipher cipher.Block
}
