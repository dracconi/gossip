package client

import (
	"bufio"
	"fmt"
	"net"

	"github.com/dracconi/gossip/handler"
	tui "github.com/marcusolsson/tui-go"
)

var srv handler.Server

func closeConn(conn net.Conn) {
	conn.Close()
}

// Loop that fetches messages
func fetchMsg(conn net.Conn, msgs *tui.List, scroll *scrollArea) {
	for {
		msg, _ := bufio.NewReader(conn).ReadString('\n')
		msgs.AddItems(msg)
		scroll.Scroll(0, -1)
	}
}

// SendMsg Function that writes messages to the server
func sendMsg(conn net.Conn, msg string) {
	conn.Write([]byte(msg + "\n"))
}

// Client compartment
func Client(ipaddr string) {
	fmt.Println("works client async")
	srv.Conn, _ = net.Dial("tcp", ipaddr)

	srv.Conn.Write([]byte("Hello! \n"))

	renderUI(ipaddr)

	// go fetchMsgs(srv.Conn)

	// for {
	// 	send, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	// 	srv.Conn.Write([]byte(send))
	// }
}
