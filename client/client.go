package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

type server struct {
	conn net.Conn
}

var srv server

func fetchMsgs(conn net.Conn) {
	for {
		msg, _ := bufio.NewReader(conn).ReadString('\n')

		fmt.Println(msg)
	}
}

// Client compartment
func Client(ipaddr string) {
	fmt.Println("works client async")
	srv.conn, _ = net.Dial("tcp", ipaddr)

	srv.conn.Write([]byte("Hello! \n"))

	go fetchMsgs(srv.conn)

	for {
		send, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		srv.conn.Write([]byte(send))
	}
}
